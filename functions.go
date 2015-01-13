package main

import (
	"fmt"
)

func main() {
	// This will work if 'go install' has been run
	// as the variable hello gets defined in variables.go
	// and both files are in the same package.
	// However if just running go run functions.go this will 
	// throw an error
	fmt.Println(hello)
}