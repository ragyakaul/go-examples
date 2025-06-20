package main

import (
	"fmt"
	"math/rand"
)

func main() {
	list := make([]int, 10)
	for i := 0; i < len(list); i++ {
		list[i] = rand.Intn(101)
	}

	for _, v := range list {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Divisible by 2 and 3:", v)
		case v%2 == 0:
			fmt.Println("Dividible by 2:", v)
		case v%3 == 0:
			fmt.Println("Divisible by 3:", v)
		default:
			fmt.Println("Nevermind:", v)

		}
	}
}

// MAke a list, then populate it with random numbers. After that add a switch statement to check if the number is divisible by 2 or 3.
// If it is, print the number and what it is divisible by. If not, print "Nevermind".
