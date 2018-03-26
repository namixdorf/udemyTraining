package main

import "fmt"

func main() {
	x := 0
	// annonymous function
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
Closure helps us limit the scope of variables used by multiple functions without closure, for two or more funcs to have access to the same variable, that variable would need to be package scope

annonymous function
a function without a name

func expression
assinging a func to a variable
*/
