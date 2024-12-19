package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// This range iterates over each element as it’s received from queue. Because it is terminated after receiving the 2 elements, the loop will stop after the 2nd element.
	for elem := range queue {
		fmt.Println(elem)
	}
}
