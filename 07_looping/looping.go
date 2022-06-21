package main

import "fmt"

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
	// 	fmt.Println(i)
	// 	if i%2 == 0 {
	// 		i /= 2
	// 	} else {
	// 		i = 2*i + 1
	// 	}

	// }

	// i := 0
	// for ; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// i := 0
	// for { // used like while
	// 	fmt.Println(i)
	// 	if i == 5 {
	// 		break
	// 	}
	// 	i++
	// }

	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// Loop:
	// 	for i := 1; 1 <= 3; i++ {
	// 		for j := 1; j <= 3; j++ {
	// 			fmt.Println(i * j)
	// 			if i*j >= 3 {
	// 				break Loop
	// 			}
	// 		}
	// 	}

	// s := [3]int{1, 2, 3}
	// for k, v := range s {
	// 	fmt.Println(k, v)
	// }

	s := "hello! my name is alex"
	for _, v := range s {
		fmt.Println(string(v))
	}
}

// Looping:

// For statements
// 	* simple loops
// 		*-for initializer; test; incrementer {}
// 		*-for test {}
// 		*-for {}
// 	* exiting early
// 		*-break
// 		*-continue
// 		*-labels
// 	* looping over collections
// 		*- arrays, slices, maps, strings, channels
// 		*- for k, v := range collection {}
