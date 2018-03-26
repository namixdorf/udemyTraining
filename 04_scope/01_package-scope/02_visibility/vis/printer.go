package vis

import (
	"fmt"
)

// PrintVar exposed in scope. Variable is available because it is a package level var
func PrintVar() {
	fmt.Println(yourName)
}
