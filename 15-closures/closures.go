package main

import "fmt"

// This function returns another function, defined anonymously inside
// The returned function closes over the variable i to form a closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// This function nextInt captures its own i value
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// To confirm state is unique to that function, create a new one
	newInts := intSeq()
	fmt.Println(newInts())
}
