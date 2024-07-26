package main

import "fmt"

func main() {

	// basic if/else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// only if
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// Logical operators can be used, &&, ||
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// statements can proceed conditionals, e.g. declarations, which can be used in current and subsequent branches
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
