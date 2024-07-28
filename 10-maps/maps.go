// Go's associative data type, ala dictionaries has
package main

import (
	"fmt"
	"maps"
)

func main() {

	// Initialize with following syntax, map[key-type]val-type
	m := make(map[string]int)

	// Typical setting of keys
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// Get value
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// If doesn't exist, the zero value of type is returned
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// Built-in len() returns number of key-values
	fmt.Println("len:", len(m))

	// Built-in delete() removes key-value pairs
	delete(m, "k2")
	fmt.Println("map:", m)

	// Built-in clear() to remove all key-value pairs
	m["k3"] = 13
	fmt.Println("map:", m)
	clear(m)
	fmt.Println("map:", m)

	// Second return value used to indicate if key was present, used to disambiguate the zer value return type if not present
	// Here the blank identifier, "_", is ues used to ignore the value
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	m["k2"] = 13
	fmt.Println("prs:", prs)
	_, prs = m["k2"]
	fmt.Println("prs:", prs)

	// Can declare and initialize in one line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("dcl:", n)

	// Number of utility functions in the maps package
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
