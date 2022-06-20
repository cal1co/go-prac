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
