package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Write a function with a typed alias operationType.
// Have a map which contains "add" with a + , etc
// In the end you should be able to add your operations to a slice of slices and perform them using your functions

type operationType func(int, int) (int, error)

var operations = map[string]operationType{ // package level variables
	"+": add,
	"-": subtract,
	"*": multiply,
	"/": divide,
}

func add(a int, b int) (int, error) {
	return a + b, nil
}

func subtract(a int, b int) (int, error) {
	return a - b, nil
}

func multiply(a int, b int) (int, error) {
	return a * b, nil
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func errorHandling(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("No error")
	}
}

func main() {
	instructions := [][]string{}
	instructions = append(instructions, []string{"5", "+", "3"})
	instructions = append(instructions, []string{"10", "-", "4"})
	instructions = append(instructions, []string{"2", "*", "3"})
	for _, instruction := range instructions {
		// Ensure the instruction has at least 3 elements
		if numArgs := len(instruction); numArgs != 3 {
			fmt.Printf("Invalid instruction: %v, expected 3 arguments but got %d\n", instruction, numArgs)
		}

		arg1, err := strconv.Atoi(instruction[0])
		if err != nil {
			errorHandling(err)
		}

		opFunc, exists := operations[instruction[1]]
		if !exists {
			errorHandling(errors.New("Operation not found in operations map" + instruction[1]))
		}

		arg2, err := strconv.Atoi(instruction[2])
		if err != nil {
			errorHandling(err)
		}
		result, err := opFunc(arg1, arg2)
		if err != nil {
			errorHandling(err)
		}
		fmt.Printf("Result of %s %s %s = %d\n", instruction[0], instruction[1], instruction[2], result)

	}
}
