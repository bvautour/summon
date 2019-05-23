// +linux
package main

import (
	"fmt"
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

// PlatformUname ... Linux Specific implementation
func PlatformUname(meta []string) *Uname {
	return &Uname{meta[0], meta[1], meta[2]}
}
