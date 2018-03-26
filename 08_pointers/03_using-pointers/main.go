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

	*b = 42        // b says, "the value at this addres, change it to 42"
	fmt.Println(a) // 42
}

/*
thi si useful
we can pass a memroy address instead ofa bunch of values (we can pass a reference)
and then we can still change the value of whatever is stored at that memory address
this makes our programs more performant
we don't have to pass around large amounts of data

everything is PASS BY VALUE in go, btx
when we pass a memory address, we are passing a value
*/
