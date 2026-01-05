//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type TimeStruct struct {
	hours   int
	minutes int
	seconds int
}

func Parse(s string) (TimeStruct, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 3 {
		return TimeStruct{}, errors.New(fmt.Sprintf("Could not part %v into time", s))
	}
	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return TimeStruct{}, errors.New(fmt.Sprintf("Cannot convert %v to hours", parts[0]))
	}
	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return TimeStruct{}, errors.New(fmt.Sprintf("Cannot convert %v to minutes", parts[1]))
	}
	seconds, err := strconv.Atoi(parts[2])
	if err != nil {
		return TimeStruct{}, errors.New(fmt.Sprintf("Cannot convert %v to seconds", parts[2]))
	}
	return TimeStruct{hours, minutes, seconds}, nil
}
