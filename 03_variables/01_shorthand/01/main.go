package main

import (
	"fmt"
)

func main() {
	// Shorthand that infers variable types
	// this declares, initializes, inferes
	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
}
