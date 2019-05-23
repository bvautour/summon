// +darwin

package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// Uname ... Darwin
type Uname struct {
	kernel  string
	release string
	machine string
	dm      *DarwinMeta
}

func (u *Uname) String() string {
	return fmt.Sprintf("&Uname{kernel: %s, release: %s, machine: %s, dm: %v}", u.kernel, u.release, u.machine, u.dm)
}

// PlatformUname ... Darwin Specific implementation
func PlatformUname(meta []string) *Uname {
	return &Uname{meta[0], meta[1], meta[2], NewDarwinMeta()}
}

// DarwinMeta ...
type DarwinMeta struct {
	osxName    string
	osxVersion string
	osxBuild   string
}

func processDarwinMeta(b []byte) *DarwinMeta {
	s := string(b)
	proc := strings.Split(s, "\n")
	dm := DarwinMeta{}
	for i, v := range proc {
		switch v {
		case "ProductName":
			dm.osxName = proc[i+1]
		case "ProductVersion":
			dm.osxVersion = proc[i+1]
		case "ProductBuildVersion":
			dm.osxBuild = proc[i+1]
		}
	}
	return &dm
}

// NewDarwinMeta ...
func NewDarwinMeta() *DarwinMeta {
	b, err := exec.Command("awk", "-F", "<|>", "/key|string/ {print $3}", "/System/Library/CoreServices/SystemVersion.plist").Output()
	if err != nil {
		log.Fatal(err)
	}
	return processDarwinMeta(b)
}

func (dm *DarwinMeta) String() string {
	return fmt.Sprintf("&DarwinMeta{osxName: %s, osxVersion: %s, osxBuild: %s}", dm.osxName, dm.osxVersion, dm.osxBuild)
}
