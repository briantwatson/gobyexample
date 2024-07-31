package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("Write ", i, " as ")
	// Basic switch
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	// You can use commas to separate multiple expressions
	case time.Saturday, time.Sunday:
		fmt.Println("I'm doing this on a weekend")
	default: // optional default
		fmt.Println("It's a weekday")
	}

	t := time.Now()

	// Switch without an expression is an alternate way to express if/else
	// Also case statements can be non-constants
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// Type switch compares types instead of values.  You can use this to discover the type of an interface value
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
