//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	commandCount := 0
	nonBlankLines := 0

	for {
		input, inputErr := reader.ReadString('\n')
		word := strings.TrimSpace(input)
		if word == "" {
			continue
		}
		if strings.ToLower(word) == "q" {
			fmt.Println("Exiting...")
			break
		}
		if strings.ToLower(word) == "bye" {
			fmt.Println("Take care!")
			commandCount++
		}
		if strings.ToLower(word) == "hello" {
			fmt.Println("Hello! How are you today?")
			commandCount++
		}
		if inputErr != nil {
			fmt.Printf("Error %v", inputErr)
			break
		}
		if inputErr == io.EOF {
			fmt.Println("Terminating")
			break
		}
		nonBlankLines++
	}
}
