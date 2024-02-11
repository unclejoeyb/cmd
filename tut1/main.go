package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello, World!"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 0
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("The result of the integer division is %v with a remainder of %v", result, remainder)
}

func printMe(printValue string) {
	println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("division by zero is not allowed")
		return 0, 0, err
	
	}
	var result = numerator / denominator
	var remainder = numerator%denominator
	return result, remainder, err
}
