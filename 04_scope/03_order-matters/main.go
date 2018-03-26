package main

import "fmt"

func main() {
	fmt.Println(x)
	fmt.Println(y)
	x := 42 // cann not be found
}

var y = 42 // order doesn't matter for y - it is in outer scope...so it is accessible in inner scopes
