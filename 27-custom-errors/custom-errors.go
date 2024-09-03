// It's possible to use custom types as errors by implementing the Error() method on them
package main

import (
	"errors"
	"fmt"
)

// It's idiomatic to suffix the name of the error type with "Error"
type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	_, err := f(42)
	var ae *argError

	// errors.As() is more advanced version of errors.Is(). It checks that a given error, or any in its chain, matches a specific error type and converts to a value of that type, returning true if so.
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err is not of type *argError")
	}
}
