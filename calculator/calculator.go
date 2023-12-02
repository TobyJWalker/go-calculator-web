package calculator

import (
	"strconv"
	"strings"
)

// Equation struct
type Equation struct {
	Num1 float64
	Num2 float64
	Op string
}

// parse user input
func ParseInput(input string) (*Equation, bool) {

	// split input string
	input_arr := strings.Split(input, " ")

	// check for valid input
	if len(input_arr) != 3 {
		return nil, true
	}

	// extract values
	num1, err_1 := strconv.ParseFloat(input_arr[0], 64) // float64
	op := input_arr[1]
	num2, err_2 := strconv.ParseFloat(input_arr[2], 64) // float64

	// check for valid operator
	if op != "+" && op != "-" && op != "*" && op != "/" {
		return nil, true
	}

	// check for valid numbers
	if err_1 != nil || err_2 != nil {
		return nil, true
	}

	// create equation
	eq := &Equation{
		Num1: num1,
		Num2: num2,
		Op: op,
	}

	// return equation, no errors
	return eq, false

}