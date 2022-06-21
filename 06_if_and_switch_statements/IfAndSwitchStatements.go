package main

import (
	"fmt"
)

func main() {
	// if true {
	// 	fmt.Println("This test is true")
	// }

	// nameAge := map[string]int{
	// 	"Alex": 21,
	// 	"Ro":   26,
	// }
	// if pop, ok := nameAge["Ro"]; ok {
	// 	fmt.Println(pop)
	// }

	// number := 50
	// guess := 70

	// if guess < number {
	// 	fmt.Println("Too Low")
	// }

	// if guess > number {
	// 	fmt.Println("Too High")
	// }

	// if guess == number {
	// 	fmt.Println("You got it!")
	// }
	// fmt.Println(number <= guess, number >= guess, number != guess)

	// switch i := 2 + 3; i {
	// case 1, 5, 10:
	// 	fmt.Println("one, five, or ten")
	// case 2, 4, 6:
	// 	fmt.Println("two")
	// default:
	// 	fmt.Println("another number")
	// }

	// i := 10
	// switch {
	// case i <= 10:
	// 	fmt.Println("less than or equal to ten")
	// 	// fallthrough
	// case i <= 20:
	// 	fmt.Println("less than or equal to twenty")
	// default:
	// 	fmt.Println("greater than twenty")
	// }

	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
		break
		fmt.Println("This will print too")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("i is another type")
	}
}

// If and Switch Statements:

// -- if statements --
// Initializer
// Comparison operators
// Logical Operators
// Short circuiting
// If - else statements
// if - else if statements
// Equality and floats

// -- Switch Statements --
// Switching on a tag
// Cases with multiple tests
// Initializers
// Switches with no tag
// Fallthrough
// Type switches
// Breaking out early
