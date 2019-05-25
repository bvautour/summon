package uname

import (
	"regexp"
	"runtime"
	"strings"
	"testing"
)

//Linux 3.10.0-514.26.2.el7.x86_64 x86_64
func TestNewUname(t *testing.T) {
	u, _ := NewUname()
	// check u.Kernel
	os := strings.Title(runtime.GOOS)
	if u.Kernel != os {
		t.Errorf("Kernel is not %v, got %v", u.Kernel, os)
	}
	// check u.Release
	semver, _ := regexp.Compile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(-[a-zA-Z\d][-a-zA-Z.\d]*)?(\+[a-zA-Z\d][-a-zA-Z.\d]*)?$`)
	if !semver.MatchString(u.Release) {
		t.Errorf("Failed to match Release: %v with semver regex", u.Release)
	}
	// Check u.Machine
	machine, _ := regexp.Compile(`\w+`)
	if !machine.MatchString(u.Machine) {
		t.Errorf(" Failed to match Machine: %v with machine regex", u.Machine)
	}
}

func TestProcessUname(t *testing.T) {
	test := processUname([]byte("Darwin 17.7.0 x86_64\n"))
	lastLetter := test[len(test)-1][len(test[len(test)-1])-1]
	if lastLetter == '\n' {
		t.Errorf("Last string should not contain newline at the end. Got: '%v'", lastLetter)
	}
}
