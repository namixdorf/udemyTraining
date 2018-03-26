package main

import (
	"fmt"
)

// iota increments at a block level
const (
	A = iota // 0
	B        // 1
	C        // 2
)

// iota increments at a block level
const (
	D = iota // 0
	E        // 1
	F        // 2
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
}
