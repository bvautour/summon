package main

import (
	"fmt"

	"github.com/bvautour/summon/pkg/hue"
)

func main() {
	m := CoreGen()
	fmt.Println(m)
	fmt.Println(hue.Red("Hello World"))
	fmt.Println(hue.Black("Hello World"))
	fmt.Println(hue.Green("Hello World"))
	fmt.Println(hue.Blue("Hello World"))
	fmt.Println(hue.Cyan("Hello World"))
	fmt.Println(hue.Yellow("Hello World"))
	fmt.Println(hue.Magenta("Hello World"))
}
