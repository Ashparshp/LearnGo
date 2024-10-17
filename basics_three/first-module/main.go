package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	color.Cyan("Hello, World!")

	str := color.BlueString("Hello, World!")
	fmt.Println(str)
}