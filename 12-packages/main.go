package main

import (
	"fmt"

	tempconv "github.com/ragyakaul/go-examples/12-packages/tempconv"
)

func main() {
	k := tempconv.Kelvin(10.0)
	c := tempconv.KelvinToCelsius(k)
	fmt.Println(c)
}
