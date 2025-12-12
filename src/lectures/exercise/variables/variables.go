//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	//Requirements:
	//* Store your favorite color in a variable using the `var` keyword
	var favouriteColour = "red"
	fmt.Println("my fav colour is", favouriteColour)

	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	var year, age = 1983, 42
	fmt.Println("my age is", age, "and I was born in", year)

	//* Store your first & last initials in two variables using block assignment
	var (
		firstInitial = 'D'
		lastInitial  = 'H'
	)
	fmt.Println("my first initial is", firstInitial, "and last initial is", lastInitial)
	//* Declare (but don't assign!) a variable for your age (in days),
	//  then assign it on the next line by multiplying 365 with the age
	// 	variable created earlier
	var ageInDays int
	ageInDays = age * 365
	fmt.Println("my age in days is", ageInDays)
}
