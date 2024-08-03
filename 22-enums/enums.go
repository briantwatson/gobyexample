// Enumerated types are a special case of sum types.  Sum types:  https://en.wikipedia.org/wiki/Algebraic_data_type
// Go doesn't have an enum type as a language feature, but enums are implemented using existing language idioms
package main

// Our enum type ServerState has underlying int type
type ServerState int

// Possible values of ServerState defined as constants
const (
	StateIdle = iota // iota keyword generates successive constant values automatically
	StateConnected
	StateError
	StateRetrying
)

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
