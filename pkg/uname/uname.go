package uname

import (
	"os/exec"
	"strings"
)

func processUname(b []byte) []string {
	s := strings.TrimRight(string(b), "\n")
	return strings.Split(s, " ")
}

// NewUname builds a new Uname struct from an exec call to uname
func NewUname() (*Uname, error) {
	b, err := exec.Command("uname", "-srm").Output()
	if err != nil {
		return nil, err
	}
	return PlatformUname(processUname(b))
}
