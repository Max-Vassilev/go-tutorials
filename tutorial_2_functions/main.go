package main

import (
	"errors"
	"fmt"
)

func main() {
	var numerator int = 12
	var denominator int = 3
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("Result: %v.\n", result)
	} else {
		fmt.Printf("Result: %v. Remainder: %v\n", result, remainder)
	}
}

func intDivision(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, nil
}
