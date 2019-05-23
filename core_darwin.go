// +build darwin

package main

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// Uname ...
type Uname struct {
	kernel  string
	release string
	machine string
}

func (u *Uname) String() string {
	return fmt.Sprintf("&Uname{kernel: %s, release: %s, machine: %s}", u.kernel, u.release, u.machine)
}

func processUname(b []byte) ([]string, error) {
	s := string(b)
	s = strings.TrimRight(s, "\r\n")
	out := strings.Split(s, " ")
	if len(out) != 3 {
		return nil, errors.New("Uname: length of uname output was not 3")
	}
	return out, nil
}

// NewUname ...
func NewUname() *Uname {
	// handles Darwin, write something to handle nix.
	b, err := exec.Command("uname", "-srm").Output()
	if err != nil {
		log.Fatal(err)
	}
	un, err := processUname(b)
	if err != nil {
		log.Fatal(err)
	}
	// check if osx again
	return &Uname{un[0], un[1], un[2]}
}

// Core ...
type Core struct {
	Uname *Uname
}

// CoreGen ...
func CoreGen() *Core {
	return &Core{
		Uname: NewUname(),
	}
}

func (c *Core) String() string {
	return fmt.Sprintf("&Core{Uname: %v}", c.Uname)
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
