package main

import "fmt"

// Variadic functions can name a variable number of arguments.
// The type of the variable argument must be specified, but the number of arguments can vary.

// Creat a function addTo that adds a base number to a variable number of arguments

func addTo(base int, numbers ...int) []int { // guess it should return the slice in the end
	result := make([]int, 0, len(numbers)) // Create a slice with capacity equal to the number of arguments
	fmt.Printf("Address Before append: %p\n", result)
	for _, number := range numbers {
		result = append(result, base+number)
		fmt.Printf("Address after append: %p\n", result) // Address stays the same because we preallocated the slice with enough capacity
	}

	return result
}

func main() {
	slice1 := []int{1, 2, 3}

	resultslice := addTo(10, slice1...) // this format unpacks the slice into individual arguments
	fmt.Println(resultslice)
}
