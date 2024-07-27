// Arrays in Go are a numbered sequence of elements of _specific_ length
package main

import "fmt"

func main() {

	// Integer array of length 5, by default zero valued
	var a [5]int
	fmt.Println("emp:", a)

	// Setting values
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// Length of an array
	fmt.Println("len:", len(a))

	// Declare and initialize in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Compiler can count number of elements with "..."
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Can specify an index with ":", elements inbetween will be zeroed
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// Can compose types to build multi-dimensionsal structures
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2di:", twoD)

	// Can create and initialize at once as well
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2di:", twoD)
}
