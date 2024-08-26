package main

import (
	"errors"
	"fmt"
)

// In Go, it's idiomatic to communicate errors via an explicit, separate return value.
// By convention, errors are the last return value and have type error, a built-in interface.
func f(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs a basic error value with the given error message.
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

var ErrorOutOfTea = fmt.Errorf("out of tea")
var ErrPower = fmt.Errorf("out of power")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrorOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("out of power: %w", ErrPower)
	}
	return nil
}

func main() {

	for _, i := range []int{7, 42} {
		// It's common to check for an error value in a single if statement.
		// So first executes the function, assigns the result to r and the error to e, then checks if e is not nil.
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	// errors.Is checks if any error in the chain matches the target error.  Useful for wrapped or nested errors.
	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrorOutOfTea) {
				fmt.Println("Out of tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Out of power!")
			} else {
				fmt.Printf("Unknown error: %s\n", err)
			}
			continue
		}
		fmt.Println("Tea is ready!")
	}
}
