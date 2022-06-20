package main

import (
	"fmt"
)

const (
	_ = iota
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func main() {
	var specialistType int
	fmt.Printf("%v\n", specialistType == catSpecialist)
}

// Constants:

// Immutable, but can be shadows.

// Replaced by the compiler at compile time
// 	* Value must be calculable at compile time
// Named like variables
// 	* PascalCase for exported constants
// 	* camelCase for internal constants

// Typed constants work like immutable variables
// 	* Can interpolate only same string
// untyped constants work like literals
// 	* Can interpolate with similar types

// Enumerated constants
// 	* Special symbol iota allows related constants to be created easily
// Iota starts at 0 in each const block and increments by one
// Watch out of constant values that match zero values for variables

// Enumerated expressions
// 	* Operations that can be determined at compile time are allowed
// 	*-Arithmetic
// 	*-Bitwise operations
// 	*-Bitshifting
