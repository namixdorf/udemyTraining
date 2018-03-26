package main

import (
	"fmt"
)

// iota increments at a block level
const (
	A = iota // 0
	B = iota // 1
	C = iota // 2
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
