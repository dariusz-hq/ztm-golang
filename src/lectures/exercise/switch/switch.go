//--Summary:
//  Create a program to display a classification based on age.
//

package main

import "fmt"

func main() {
	//--Requirements:
	//* Use a `switch` statement to print the following:
	//  - "newborn" when age is 0
	//  - "toddler" when age is 1, 2, or 3
	//  - "child" when age is 4 through 12
	//  - "teenager" when age is 13 through 17
	//  - "adult" when age is 18+

	age := 15

	switch true {
	case age == 0:

		fmt.Println("newborn")
	case age > 0 && age < 4:
		fmt.Println("toddler")
	case age > 3 && age < 13:
		fmt.Println("child")
	case age > 12 && age < 18:
		fmt.Println("teenager")
	default:
		fmt.Println("adult")
	}
}
