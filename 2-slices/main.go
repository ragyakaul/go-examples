package main

import (
	"fmt"
)

func main() {
	// Create slice of length 3 and capacity 3
	slice1 := make([]int, 3, 3)
	slice1[0] = 1
	slice1[1] = 4
	slice1[2] = 5

	fmt.Println("Slice 1:", slice1)
	fmt.Printf("Slice 1 address before calling modifySlice: %p\n", slice1) // Prints the address of slice1
	modifySlice(slice1)

	// Shallow copy
	slice2 := slice1
	fmt.Println("Slice 2 shallow copy from slice 1:", slice2)

	// Deep copy
	slice3 := make([]int, len(slice1))
	copy(slice3, slice1)

	slice2[0] = 3

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

}

func modifySlice(slice1 []int) {

	slice1[1] = 1
	fmt.Printf("Address Before append: %p\n", slice1)
	fmt.Println(slice1)
	slice1 = append(slice1, 6)
	fmt.Printf("Address After append: %p\n", slice1)
	slice1[2] = 2
	fmt.Println(slice1)
}

//
