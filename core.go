package main

import (
	"fmt"

	"github.com/bvautour/summon/pkg/uname"
)

// Core ...
type Core struct {
	Uname *uname.Uname
}

// CoreGen ...
func CoreGen() *Core {
	return &Core{
		Uname: uname.NewUname(),
	}
}

func (c *Core) String() string {
	return fmt.Sprintf("&Core{Uname: %v}", c.Uname)
}
