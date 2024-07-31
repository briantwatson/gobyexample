// Structs are typed collections of fields.  Useful for grouping data to form records
package main

import (
	"fmt"
	"strconv"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	// Can safely return a pointer to local variable. Local variables will survive scope of function
	return &p
}

// Experiment with structs and 'static-like' variables
type bicycle struct {
	brand        string
	model        string
	serialNumber string
}

var serialCounter int

func newBicycle(model string) *bicycle {
	serialCounter++
	b := bicycle{brand: "Specialized", model: model, serialNumber: "SN" + strconv.Itoa(serialCounter)}
	return &b
}

func main() {
	// CREATION
	// --------
	// Create a new struct in function call implicitly defining fields based on order
	fmt.Println(person{"Bob", 20})

	// Can explicitly define fields
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields are zero-valued
	fmt.Println(person{name: "Fred"})

	// & prefix yields pointer to struct
	fmt.Println(&person{name: "Ann", age: 40})

	// It's _idiomatic_ to encapsulate struct creation in constructor functions
	fmt.Println(newPerson("Jon"))

	// Access properties with dot notation
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s            // struct pointer
	fmt.Println(sp.age) // can use dot notation with pointers, automatically dereferenced

	// Structs are mutable!
	sp.age = 51
	fmt.Println(sp.age)

	// If only used once, structs don't need a named type, can have anonymous type
	// Commonly used for table-driven tests:  https://gobyexample.com/testing-and-benchmarking
	dog := struct {
		name   string
		isGood bool
	}{"Silas", true}
	fmt.Println(dog)

	bike1 := newBicycle("Allez")
	bike2 := newBicycle("Tarmac")
	bike3 := newBicycle("Stumpjumper")

	fmt.Println(bike1)  // prints pointer, which also dereferences to print value
	fmt.Println(*bike2) // dereferencing for comparison
	fmt.Println(*bike3)
}
