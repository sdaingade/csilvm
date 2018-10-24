package csilvm

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestFailure1Of1DeviceLVLinear(t *testing.T) {
	prepareSCSIDebug()
	defer clearSCSIDebug()
	scsivols := listSCSIDebugDevices()
	vgname := testvgname()
	pvname := scsivols[0]
	client, clean := startTest(vgname, []string{pvname})
	defer clean()
	// Create the volume that we'll be publishing.
	createReq := testCreateVolumeRequest()
	createResp, err := client.CreateVolume(context.Background(), createReq)
	if err != nil {
		t.Fatal(err)
	}
	volumeId := createResp.GetVolume().GetId()
	// Prepare a temporary mount directory.
	tmpdirPath, err := ioutil.TempDir("", "csilvm_tests")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdirPath)
	targetPath := filepath.Join(tmpdirPath, volumeId)
	if err := os.Mkdir(targetPath, 0755); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(targetPath)
	// Publish the volume to /the/tmp/dir/volume-id
	publishReq := testNodePublishVolumeRequest(volumeId, targetPath, "xfs", nil)
	_, err = client.NodePublishVolume(context.Background(), publishReq)
	if err != nil {
		t.Fatal(err)
	}
	// Remove the underlying devmapper mapping when the test is completed.
	defer forceRemoveVolume(targetPath, vgname, volumeId)

	// Remove the loop device underlying the PV.
	log.Printf("Removing SCSI device")
	removeSCSIDevice(pvname)
	log.Printf("Removed SCSI device")
	// Assert that Probe fails.
	_, err = client.Probe(context.Background(), testProbeRequest())
	exp := status.Errorf(
		codes.FailedPrecondition,
		fmt.Sprintf("Cannot lookup physical volume %q: err=lvm: physical volume not found", pvname))
	if !grpcErrorEqual(err, exp) {
		t.Error(err.Error())
	}
}

func prepareSCSIDebug() {
	log.Print("preparing and loading scsi_debug module")
	out, err := exec.Command("lsmod").CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("prepareSCSIDebug: cannot run `lsmod`: err=%v: %s", err, string(out)))
	}
	moduleLoaded := bytes.Contains(out, []byte("scsi_debug"))
	// If the module is loaded, unload it.
	if moduleLoaded {
		if err := exec.Command("rmmod", "scsi_debug").Run(); err != nil {
			panic(err)
		}
	}
	// Load the scsi_debug module with 4 devices, each of size `pvsize`.
	pvsizeMiB := pvsize >> 20
	// See http://sg.danny.cz/sg/sdebug26.html
	options := []string{
		fmt.Sprintf("dev_size_mb=%d", pvsizeMiB),
		"max_luns=1", // number of devices
		"num_tgts=1", // number of targets
	}
	cmd := exec.Command("modprobe", append([]string{"scsi_debug"}, options...)...)
	log.Print(cmd)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func clearSCSIDebug() {
	log.Print("rmmod scsi_debug")
	out, err := exec.Command("rmmod", "scsi_debug").CombinedOutput()
	if err != nil {
		log.Print(string(out))
		panic(err)
	}
}

/*
// This is example output of having multiple regular as well as scsi_debug hosts:
]# lsscsi -H
[0]    ahci
[1]    ahci
[2]    scsi_debug
[3]    scsi_debug
]# lsscsi
[2:0:0:0]    disk    Linux    scsi_debug       0187  /dev/sda
[3:0:0:0]    disk    Linux    scsi_debug       0187  /dev/sdb
]# lsscsi 2
[2:0:0:0]    disk    Linux    scsi_debug       0187  /dev/sda
*/
func listSCSIDebugDevices() []string {
	// Find scsi_debug host address
	host := ""
	out, err := exec.Command("lsscsi", "-H").CombinedOutput()
	for _, line := range bytes.Split(out, []byte("\n")) {
		if !bytes.Contains(line, []byte("scsi_debug")) {
			continue
		}
		fields := bytes.Fields(line)
		if len(fields) != 2 {
			panic(fmt.Errorf("listSCSIDebugDevices: cannot parse lsscsi -H: %s", line))
		}
		line = bytes.TrimPrefix(fields[0], []byte("["))
		line = bytes.TrimSuffix(line, []byte("]"))
		host = string(line)
	}
	// host now refers to the last scsi_debug in the list
	if host == "" {
		panic(fmt.Errorf("listSCSIDebugDevices: cannot find scsi_debug host, perhaps scsi_debug module is not loaded"))
	}
	// Find scsi devices that belong to the selected scsi_debug host.
	out, err = exec.Command("lsscsi", host).CombinedOutput()
	if err != nil {
		panic(err)
	}
	var devices []string
	for _, line := range bytes.Split(out, []byte("\n")) {
		if !bytes.Contains(line, []byte("scsi_debug")) {
			continue
		}
		fields := bytes.Fields(line)
		if len(fields) != 6 {
			panic(fmt.Errorf("listSCSIDebugDevices: cannot parse lsscsi %s: %s", host, line))
		}
		sg := fields[5]
		devices = append(devices, string(sg))
	}
	return devices
}

func removeSCSIDevice(name string) {
	log.Print("removing scsi device", name)
	out, err := exec.Command("sh", "-c", fmt.Sprintf("echo 1 > /sys/class/block/%s/device/delete", filepath.Base(name))).CombinedOutput()
	if err != nil {
		log.Print(string(out))
		panic(err)
	}
}

func forceRemoveVolume(mountpoint, vgname, lvname string) {
	log.Printf("unmounting %q", mountpoint)
	cmd := exec.Command("umount", mountpoint)
	log.Print(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Print(string(out))
		log.Print(err)
	} else {
		log.Print("unmount succeeded")
	}
	log.Print("forcefully removing logical volume")
	cmd = exec.Command("dmsetup", "remove", fmt.Sprintf("/dev/%s/%s", vgname, lvname))
	log.Print(cmd)
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Print(string(out))
		log.Print(err)
		panic(err)
	}
}
