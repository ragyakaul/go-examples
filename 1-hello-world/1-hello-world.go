package main

import (
	"fmt"

	"rsc.io/quote/v4"
)

func main() {
	const age = 42 // shouldn't we specify a type?
	const name string = "Alice"
	const unused float64 = age // aren't all technically unused? also, are we just converting types here? is this the same as `const unused = float64(age)`?

	fmt.Printf("Hello %s (%f years old or %d)\n", name, unused, age)
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
}

// Mak a const age, name and unused (which is a float64) amd copy ag into it.
// Print out name and age in formatted string
