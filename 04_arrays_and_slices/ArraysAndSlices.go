package main

import (
	"fmt"
)

func main() {
	// var students [3]string
	// fmt.Printf("Students: %v\n", students)
	// students[0] = "Lisa"
	// students[1] = "Ahmed"
	// students[2] = "Arnold"
	// fmt.Printf("Student 1: %v\n", students[1])
	// fmt.Printf("Num of students: %v\n", len(students))

	// var identityMatrix [3][3]int
	// identityMatrix[0] = [3]int{1, 0, 0}
	// identityMatrix[1] = [3]int{0, 1, 0}
	// identityMatrix[2] = [3]int{0, 0, 1}
	// fmt.Println(identityMatrix)

	// a := [...]int{1, 2, 3}
	// b := &a
	// b[1] = 5
	// fmt.Println(a)
	// fmt.Println(b)

	// a := []int{1, 2, 3}
	// b := a
	// b[1] = 5
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))

	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// b := a[:]   // slice of all elements
	// c := a[3:]  // slice from 4th element to end
	// d := a[:6]  // slice first 6 elements
	// e := a[3:6] // slice the 4th, 5th, and 6th elements
	// a[5] = 42
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	// a := make([]int, 3, 100)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))

	// a := []int{}
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))
	// a = append(a, 1)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))
	// a = append(a, 2, 3, 4, 5)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))

	// a := []int{1, 2, 3, 4, 5}
	// // b := a[1:]
	// b := a[:len(a)-1]
	// fmt.Println(b)

	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	b := append(a[:2], a[3:]...)
	fmt.Println(b)
	fmt.Println(a)
}

// Arrays & Slices:

// -- Arrays --
// * Collection of items with same type
// * Fixed size
// * Declaration styles
//	*-a := [3]int{1, 2, 3}
// 	*-a := [...]int{1, 2, 3}
// 	*-var a [3]int
// * Access via zero-based index
// 	*-a := [3]int {1, 3, 5} // a[1] == 3
// * len function returns size of array
// * Copies refer to different underlying data

// -- Slices --
// * Backed by array
// * Creation styles:
// 	*-Slices existing array or slice
// 	*-Literal style
// 	*-Via make function
// 		** a := make([]int, 10) // create slice with capacity and length == 10
// 		** a := make([]int, 10, 100) // slice with length == 10 and capacity == 100
// * len function returns length of slice
// * cap function returns length of underlying array
// * append function to add elements to slice
// 	*-May cause expensive copy operation if underlying array is too small
// * Copies refer to same underlying array.
