package main

import (
	"fmt"
)

func main() {
	nameAge := make(map[string]int)
	nameAge = map[string]int{
		"Alex": 21,
		"Ro":   26,
	}
	fmt.Println(nameAge)
}
