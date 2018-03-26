package main

import (
	"fmt"
)

func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x208102a220

	var b *int = &a
	fmt.Println(b)  // 0x208102a220
	fmt.Println(*b) // 43
}

/*
b is an int points;
b pionts to the memory address where an int is stored
to see the value in that memory address, add a * in from of b
this is known as deferencing
the * is an operator in this case
the * says give me that value of what is stored in the memory address (dereferencing)
*/
