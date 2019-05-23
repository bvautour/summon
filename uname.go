package main

import (
	"log"
	"os/exec"
	"strings"
)

func processUname(b []byte) []string {
	s := strings.TrimRight(string(b), "\r\n")
	return strings.Split(s, " ")
}

// NewUname ...
func NewUname() *Uname {
	b, err := exec.Command("uname", "-srm").Output()
	if err != nil {
		log.Fatal(err)
	}
	// run platform uname to cover linux or mac
	return PlatformUname(processUname(b))
}
