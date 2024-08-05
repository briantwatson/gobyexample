// Enumerated types are a special case of sum types.  Sum types:  https://en.wikipedia.org/wiki/Algebraic_data_type
// Go doesn't have an enum type as a language feature, but enums are implemented using existing language idioms
package main

import "fmt"

// Our enum type ServerState has underlying int type
type ServerState int

// Possible values of ServerState defined as constants
const (
	StateIdle = iota // iota keyword generates successive constant values automatically
	StateConnected
	StateError
	StateRetrying
)

// Map with possible string values.  Can be cumbersome with many.  string tool can be used:  https://pkg.go.dev/golang.org/x/tools/cmd/stringer
// Also relevant:  https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

// Implementing fmt.Stringer interface.  https://pkg.go.dev/fmt#Stringer
func (s ServerState) String() string {
	return stateName[s]
}

// Emulates a state transition for a server
func transition(s ServerState) ServerState {
	fmt.Printf("Transition called with state: %s (int value: %d)\n", s, s) // Debug statement
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)

	// Works because ServerState has an underlying int type
	ns3 := transition(3)
	fmt.Println(ns3)

	//The compiler complains about the following type mismatch, provides some compile time safety
	//var i int = 3
	// ns4 := transition(i)
	//fmt.Println(ns4)

}
