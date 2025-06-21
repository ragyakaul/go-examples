package tempconv

func CelsiusToFahrenheit(c Celsius) Fahrenheit {
	var result Celsius = c*9/5 + 32
	return Fahrenheit(result) // casts Celsius (with base type float64) to Fahrenheit
}
func FahrenheitToCelsius(f Fahrenheit) Celsius {
	var result Fahrenheit = (f - 32) * 5 / 9
	return Celsius(result)
}

func KelvinToCelsius(c Celsius) Kelvin {
	var result Celsius = c - 273.15
	return Kelvin(result)
}

func FahrenheitToRankine(f Fahrenheit) Rankine {
	var result Fahrenheit = f + 459.67
	return Rankine(result)
}
