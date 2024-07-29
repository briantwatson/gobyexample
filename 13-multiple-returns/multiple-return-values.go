package main

import "fmt"

// Go supports multiple returns and is used often in idiomatic Go
func vals() (int, int) {
	return 3, 7
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// If you want a subset, use the blank indicator
	_, c := vals()
	fmt.Println(c)
}
