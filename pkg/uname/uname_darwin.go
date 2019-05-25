// +darwin

package uname

import (
	"fmt"
	"os/exec"
	"strings"
)

// Uname ... Darwin
type Uname struct {
	Kernel  string
	Release string
	Machine string
	Dm      *DarwinMeta
}

func (u *Uname) String() string {
	return fmt.Sprintf("&Uname{Kernel: %s, Release: %s, Machine: %s, Dm: %v}", u.Kernel, u.Release, u.Machine, u.Dm)
}

// PlatformUname ... Darwin Specific implementation
func PlatformUname(meta []string) (*Uname, error) {
	dm, err := NewDarwinMeta()
	if err != nil {
		return nil, err
	}
	return &Uname{meta[0], meta[1], meta[2], dm}, nil
}

// DarwinMeta ...
type DarwinMeta struct {
	OsxName    string
	OsxVersion string
	OsxBuild   string
}

func processDarwinMeta(b []byte) *DarwinMeta {
	s := string(b)
	proc := strings.Split(s, "\n")
	dm := DarwinMeta{}
	for i, v := range proc {
		switch v {
		case "ProductName":
			dm.OsxName = proc[i+1]
		case "ProductVersion":
			dm.OsxVersion = proc[i+1]
		case "ProductBuildVersion":
			dm.OsxBuild = proc[i+1]
		}
	}
	return &dm
}

// NewDarwinMeta ...
func NewDarwinMeta() (*DarwinMeta, error) {
	b, err := exec.Command("awk", "-F", "<|>", "/key|string/ {print $3}", "/System/Library/CoreServices/SystemVersion.plist").Output()
	if err != nil {
		return nil, err
	}
	return processDarwinMeta(b), nil
}

func (dm *DarwinMeta) String() string {
	return fmt.Sprintf("&DarwinMeta{OsxName: %s, OsxVersion: %s, OsxBuild: %s}", dm.OsxName, dm.OsxVersion, dm.OsxBuild)
}
