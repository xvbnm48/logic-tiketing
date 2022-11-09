package main

import "fmt"

func main() {
	// you are making a ticketing system, for children under 3 years old, they get in for free
	// need to take the ages of 5 passengers and calculate the total cost of the tickets
	// if the passenger is under 3, they get in for free

	// create a slice of ints called ages, and assign 5 ages to it
	ages := []int{2, 5, 42, 24, 18, 10}

	// create a variable called total and assign it the value of 0
	total := 0

	// use a for range loop to iterate over the ages slice and calculate the ticket price for each

	for _, age := range ages {
		if age < 3 {
			total += 0

		} else {
			total += 100
		}
	}

	// print out the total ticket price for all 5 passengers
	fmt.Println(total)
}
