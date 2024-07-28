// Range iterates over elements in a number of data structures
package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// Range on arrays and slices provides both index and value of each element
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// Range on map iterates over key-value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Range can also just iterate over keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// Range can also iterate over values using the blank identifier
	for _, v := range kvs {
		fmt.Println("val:", v)
	}

	// Range over strings iterates over Unicode code points
	// First value is starting byte index of rune and second is rune itself
	// See https://gobyexample.com/strings-and-runes
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
