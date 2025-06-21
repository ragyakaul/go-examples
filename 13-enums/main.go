package main

import "fmt"

/*
Enums are used to represent a fixed set of related values,
improving code readability and preventing errors by
restricting a variable to a predefined set of options.
*/
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednsday
	Thursday
	Friday
	Saturday
)

func (w Weekday) String() string {
	sliceDays := []string{"Sunday", "Monday", "Tuesday", "Wednsday", "Thursday", "Friday", "Saturday"}
	return sliceDays[w]
}

func main() {
	var today Weekday = Saturday

	if today == Saturday {
		fmt.Println("yay! It's the weekend!")
	}

}
