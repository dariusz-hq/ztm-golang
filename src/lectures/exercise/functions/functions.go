//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//

package main

import "fmt"

//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func greet(name string) {
	fmt.Println("hello", name)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func message() string {
	return "some message to be returned"
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func add(n1, n2, n3 int) int {
	return n1 + n2 + n3
}

// * Write a function that returns any number
func returnNumber() int {
	return 9
}

// * Write a function that returns any two numbers
func returnTwoNumbers() (int, int) {
	return 1, 2
}

func main() {
	//--Requirements:
	//* Call every function at least once
	greet("John Doe")
	fmt.Println(message())
	fmt.Println(returnNumber())
	fmt.Println(returnTwoNumbers())

	//* Add three numbers together using any combination of the existing functions.
	//  * Print the result
	fmt.Println(add(1, 2, 3))
}
