// Take a variable length of arguments. For example, fmt.Println() is a variadic function
package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sum(1, 2)
	sum(1, 2, 3)

	// If you have multiple elements in a slice, you can apply that slice to a variadic function
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
