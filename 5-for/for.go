// For is Go's only looping construct
package main

import "fmt"

func main() {

	// Most basic, one condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic initial/condition,after for loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// loop over a range N times
	for i := range 3 {
		fmt.Println("range", i)
	}

	// for without a condition-- will look until a `break` or `return`
	for {
		fmt.Println("no condition loop")
		break
	}

	// also can `continue` to next iteration of loop
	// here we should only print odds
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// Can use the following pattern to reference index and values of certain data structures
	nums := []int{2, 3, 4}
	for idx, num := range nums {
		fmt.Println("idx:", idx)
		fmt.Println("num:", num)
	}
}
