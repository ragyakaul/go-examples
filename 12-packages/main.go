package main

import (
	"fmt"

	tempconv "github.com/ragyakaul/go-basic-examples/12-packags/tempconv"
)

func main() {
	k := tempconv.Kelvin(10.0)
	c := tempconv.KelvinToCelsius(k)
	fmt.Println(c)
}
