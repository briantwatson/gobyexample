package main

import "fmt"

// Function that takes a map and returns a slice of keys.  K, are of type 'comparable'.  Values of type 'any'.
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// Example of a generic singly linked list, values of type 'any'.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elements []T
	for e := lst.head; e != nil; e = e.next {
		elements = append(elements, e.val)
	}
	return elements
}

func main() {

	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	// When invoking generic functions we can rely on type inference
	fmt.Println("keys:", MapKeys(m))

	// Can also specify types explicitly
	_ = MapKeys[int, string](m)

	lst := List[string]{}
	lst.Push("one")
	lst.Push("two")
	lst.Push("three")
	fmt.Println("list:", lst.GetAll())

	iLst := IntList{}
	iLst.Push(1)
	iLst.Push(2)
	iLst.Push(3)
	fmt.Println("iList:", iLst.GetAll())
}

// Regular type for comparison
type IntList struct {
	head, tail *IntElement
}

type IntElement struct {
	next *IntElement
	val  int
}

func (lst *IntList) Push(v int) {
	if lst.tail == nil {
		lst.head = &IntElement{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &IntElement{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *IntList) GetAll() []int {
	var elements []int
	for e := lst.head; e != nil; e = e.next {
		elements = append(elements, e.val)
	}
	return elements
}
