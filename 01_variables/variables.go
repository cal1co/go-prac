package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var i int
	// i = 42
	// var j int = 27
	// k := 99
	// fmt.Println(i, j, k)
	// fmt.Printf("%v, %T", j, j)

	// var i float32 = 42.5
	// fmt.Printf("%v, %T\n", i, i)

	// var j int
	// j = int(i) // go won't risk losing information. need an explicit conversion of variable type
	// fmt.Printf("%v, %T\n", j, j)

	var i int = 42.
	fmt.Printf("%v, %T\n", i, i)

	var j string
	j = strconv.Itoa(i) // go won't risk losing information. need an explicit conversion of variable type
	fmt.Printf("%v, %T\n", j, j)
}

// Variables

// - Variable declaration:
// var foo int
// var foo int = 42
// foo:=42

// - Can't reduce variables, but can shadow them

// - All variables must be used

// - Visibility:
// Lower case first letter for package scope
// Upper case first letter to export
// No private scope

// - Naming conventions:
// Pascal or cameCase
// 	Capitalize acronyms (HTTP, URL)
// As short as reasonable
// 	Longer names for longer lives (more uses)

// - Type conversions:
// DestinationType(variable)
// use strconv package for strings
