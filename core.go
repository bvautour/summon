package main

import (
	"fmt"
	"log"

	"github.com/bvautour/summon/pkg/uname"
)

// Core ...
type Core struct {
	Uname *uname.Uname
}

// CoreGen ...
func CoreGen() *Core {
	un, err := uname.NewUname()
	if err != nil {
		log.Fatal(err)
	}
	return &Core{
		Uname: un,
	}
}

func (c *Core) String() string {
	return fmt.Sprintf("&Core{Uname: %v}", c.Uname)
}
