package main

import "fmt"

func main() {
	// In Go, variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.
	var a = "initial"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables.  Here it's a boolean.
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
	var e int
	fmt.Println(e)

	// Here the zero value for a string is an empty string.
	var f string
	fmt.Println(f)

	// The zero value for a float is 0.
	var g float64
	fmt.Println(g)

	// The zero value for a boolean is false.
	var h bool
	fmt.Println(h)

	// The := syntax is shorthand for declaring and initializing a variable, e.g. for var i string = "apple" in this case.
	i := "apple"
	fmt.Println(i)
}
