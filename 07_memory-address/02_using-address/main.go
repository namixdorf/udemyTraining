package main

import (
	"fmt"
)

const metersToYards float64 = 1.09361

func main() {
	var meters float64 // create memory address
	fmt.Printf("Enter meters swam: ")
	fmt.Scan(&meters) // get answer and putting it into that memory address
	yards := meters * metersToYards
	fmt.Println(meters, "meters is", yards, "yards.")
}
