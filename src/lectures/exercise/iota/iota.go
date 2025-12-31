//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

const (
	Addition = iota
	Subtraction
	Multiplication
	Division
)

type Operation int

type Calculator struct {
	operation Operation
}

func (c *Calculator) calculate(n1, n2 int) int {
	switch c.operation {
	case Addition:
		return n1 + n2
	case Subtraction:

		return n1 - n2
	case Multiplication:
		return n1 * n2
	case Division:
		return n1 / n2
	default:
		panic("wrong operation")
	}
}

func main() {
	add := Calculator{Addition}
	fmt.Println(add.calculate(2, 2)) // = 4

	sub := Calculator{Subtraction}
	fmt.Println(sub.calculate(10, 3)) // = 7

	mul := Calculator{Multiplication}
	fmt.Println(mul.calculate(3, 3)) // = 9

	div := Calculator{Division}
	fmt.Println(div.calculate(100, 2)) // = 50
}
