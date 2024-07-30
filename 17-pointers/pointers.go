package main

import "fmt"

// With an int parameter, arguments are passed by value.  This function will get a copy of distinct from calling function
func zeroval(ival int) {
	ival = 0
}

// With a *int parameter, taking an int pointer.
func zeroptr(iptr *int) {
	//*iptr dereferences pointer from its memory address to current value at that address
	// Assigning a value to dereferenced pointer changes value at that address
	*iptr = 0
}

func main() {

	i := 1
	fmt.Println("initial:", i)

	// Shows we are not actually changing i in zeroval
	zeroval(i)
	fmt.Println("zeroval:", i)

	// Providing memory address of i with &var syntax
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Pointers can be printed
	fmt.Println("pointer:", &i)

}
