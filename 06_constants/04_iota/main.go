package main

import (
	"fmt"
)

// iota increments at a block level - other way to define iota
const (
	A = iota // 0
	B        // 1
	C        // 2
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
