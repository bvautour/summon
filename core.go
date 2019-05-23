package main

import (
	"fmt"
)

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
