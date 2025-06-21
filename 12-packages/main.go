package main

import (
	"fmt"

	"tempconv"
)

func main() {
	k := tempconv.Kelvin(10.0)
	c := tempconv.KelvinToCelsius(k)
	fmt.Println(c)
}
