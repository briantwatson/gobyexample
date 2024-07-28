// Go supports constants of: character, string, boolean, and numeric values
package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	// Constant expressions perform arithmetic with arbitrary precision
	const d = 3e20 / n

	// Numeric constant has no type until it's given one
	fmt.Println(int64(d))

	// Number can be given a type by using it a context that requires one, such as variable assignment or function call
	// Here math.Sin expects a float64
	fmt.Println(math.Sin(n))
}
