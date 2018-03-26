package main

import (
	"fmt"
)

// iota increments at a block level
const (
	_ = iota      // 0
	B = iota * 10 // 10
	C = iota * 10 // 20
)

func main() {
	fmt.Println(B)
	fmt.Println(C)
}
