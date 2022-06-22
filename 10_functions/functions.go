package main

import (
	"fmt"
)

// func main() {
// 	// sayMessage("Hello, Go!")
// 	// for i := 0; i < 5; i++ {
// 	// 	sayMessage("Hello Go!", i)
// 	// }

// 	// sayMessage("Hello", "Ro")

// 	// sum(1, 2, 3, 4, 5)

// 	// s := sum(1, 2, 3, 4, 5)
// 	// fmt.Println("The sum is ", s)

// 	d, err := divide(5.0, 3.0)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(d)
// }

// func sayMessage(greeting, name string) {
// 	fmt.Println(greeting, name)
// 	name = "Alex"
// 	fmt.Println(name)
// }

// func sum(values ...int) *int {
// 	// fmt.Println(values)
// 	// result := 0
// 	// for _, v := range values {
// 	// 	result += v
// 	// }
// 	// fmt.Println("The sum is ", result)

// 	fmt.Println(values)
// 	result := 0
// 	for _, v := range values {
// 		result += v
// 	}
// 	return &result
// }

// func sum(values ...int) (result int) {
// 	fmt.Println(values)
// 	for _, v := range values {
// 		result += v
// 	}
// 	return
// }

// func divide(a, b float64) (float64, error) {
// 	if b == 0.0 {
// 		return 0.0, fmt.Errorf("Cannot divide by zero")
// 	}
// 	return a / b, nil
// }

func main() {
	// for i := 0; i < 5; i++ {
	// 	func(counter int) {
	// 		fmt.Println(counter)
	// 	}(i)
	// }

	// f := func() {
	// 	fmt.Println("Hello Go!")
	// }
	// f()

	// var divide func(float64, float64) float64
	// divide = func(a, b float64) (float64, error){
	// 	if b == 0.0 {
	// 		return 0.0, fmt.Errorf("Cannot divide by zero")
	// 	} else {
	// 		retrun a / b, nil
	// 	}
	// }
	// d, err := divide(5.0, 0.0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(d)

	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

// Functions:

// -- Basc Syntax --
// * func foo(){
//		...
// 	 }

// -- Parameters --
// * Comma delimited list of variables and types
// 	-* func foo(bar string, baz int)
// * Parameters of same type list type once
// 	-* func foo(bar, baz int)
// * When pointers are passed in, the function can change the value in the caller
// 	-* This is always true for data of slices and maps
// * Use variadic parameters to send list of same types in
// 	-* Must be last parameter
// 	-* Received as a slice
// 	-* func foo (bar string, baz ...int)

// -- Return values --
// * Single return values just list type
// 	*- func foo() int
// * Multiple return value list types surrounded by parentheses
// 	*- func foo() (int, error)
// 	*- The (result type, error) paradigm is a very common idiom
// * Can use named return values
// 	*- initializes returned variable
// 	*- Return using return keyword on its own
// * Can return addresses of local variables
// 	*- Automatically promoted from local memory (stack) to shared memory (heap)

// -- Anonymous functions --
// * Functions don't have names if they are:
// 	*- IIFE
// 	*- Assigned to a variable or passed as an argument to a function
// 		** a:= func(){
// 			...
// 			}
// 			a()

// -- Functions as types --
// * Can assign functions to variables or use as arguments and return values in functions
// * Type signature is like function signature, with no parameter names
// 	*- var f func(string, string, int) (int, error)

// -- Methods --
// * Function that executes in context of a type
// * Format
// 	*- func (g greeter) greet(){
// 		...
// 		}
// * Receiver can be value or pointer
// 	*- Value receiver gets copy of type
// 	*- Pointer receiver gets pointer to type
