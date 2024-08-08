package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// A container embeds a base
type container struct {
	base
	str string
}

func main() {

	co := container{
		base: base{num: 1},
		str:  "some name",
	}

	// Can directly access the embedded structs fields on the container
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Can also spell out full path
	fmt.Println("also num:", co.base.num)

	// Since container inherits base, the methods of base become methods of container
	fmt.Println("describe:", co.describe())

	// Embedding a struct may be used to bestow interface implementation
	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}
