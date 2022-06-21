package main

import (
	"fmt"
)

func main() {
	// a := 42
	// b := a
	// fmt.Println(a, b)
	// a = 27
	// fmt.Println(a, b)

	// var a int = 42
	// var b *int = &a
	// fmt.Println(a, *b)

	// a := [3]int{1, 2, 3}
	// b := &a[0]
	// c := &a[1]
	// fmt.Printf("%v, %p %p\n", a, b, c)

	// var ms *myStruct
	// ms = new(myStruct)
	// ms.foo = 42
	// fmt.Println(ms.foo)

	// a := [3]int{1, 2, 3}
	// // a := []int{1, 2, 3}
	// b := a
	// fmt.Println(a, b)
	// a[1] = 42
	// fmt.Println(a, b)

	a := map[string]string{"foo": "bar", "baz": "buz"}
	b := a
	fmt.Println(a, b)
	a["foo"] = "qux"
	fmt.Println(a, b)
	a["foo"] = "qux"
	fmt.Println(a, b)
}

// type myStruct struct {
// 	foo int
// }

// Pointers:

// -- Creating pointers --
// 	* Pointer types use an asterisk as a prefix to type pointed to
// 		*-int - a poitner to an integer
// 	* Use the addressof operator (&) to get address of variable

// -- Dereferencing pointers --
// 	* Dereference a pointer by preceding with an asterisk
// 	* Complex types (eg structs) are automatically dereferenced

// -- Create pointers to objects --
// 	* Can use the addressof operator (&) if value type already exists
// 		*- ms:=mystruct{foo:42}
// 		p:=&ms
// 	* Use addressof operator before initializer
// 		*-&myStruct{foo:42}
// 	* Use the new keyword
// 		*- Can't initialize fields at the same time

// -- Types with internal pointers --
// 	* All assignment operations in Go are copy operations
// 	* Slices and maps contain internal pointers, so copies point to same underlying data
