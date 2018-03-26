package main

import (
	"fmt"
)

func main() {
	// Shorthand that infers variable types
	// this declares, inferes
	// no initialization
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
}
