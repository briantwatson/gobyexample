package main

import "fmt"

// Function takes two ints and returns an int
func plus(a int, b int) int {
	return a + b
}

// Multiple consecutive parameters with same type can omit like-typed names up to the final
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
