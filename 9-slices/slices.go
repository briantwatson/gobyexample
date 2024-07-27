// Slices give a more powerful interface to sequences than arrays
package main

import (
	"fmt"
	"slices"
)

func main() {

	// Only typed by element they contain--no length.  Note:  no number in declaration
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0) // equal to nil and length of zero

	// To create an empty slice with non-zero length, use `make`
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// Set and get just like with arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// Length is same as array
	fmt.Println("len:", len(s))

	// ------------------------
	// Enhanced slice operations
	// ------------------------

	// Append: returns a slice containing one or more new values
	s = append(s, "d")
	s = append(s, "e", "f") // can append multiple
	fmt.Println("apd:", s)

	// Copy:  Copies elments into source, from target
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// "Slice" operator:  syntax `slice[low:high]` with low being inclusive, 'starting at and up to, but excluding...'
	// Returns new slice with subset
	l := s[2:5] // this gets elements s[2], s[3], and s[4]
	fmt.Println("sl1:", l)

	l = s[:5] // up to, but excluding 5
	fmt.Println("sl2:", l)

	l = s[2:] // up from and including element 2
	fmt.Println("sl3:", l)

	// Declare and initialize in one line like arrays
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slices package contains a number of utility functions for them
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Can be composed into multi-dimensional structures like arrays. Inner slices length CAN vary, UNLIKE arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1 // changing size of inner slice here
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2di:", twoD)
}
