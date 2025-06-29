package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Rankine float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g degrees Celsius", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g degrees Fahrenheit", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g degrees Kelvin", k)
}
func (r Rankine) String() string {
	return fmt.Sprintf("%g degrees Rankine", r)
}
