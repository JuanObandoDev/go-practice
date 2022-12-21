package main

// calculator using switch

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// creating scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)

	// activate scanner
	scanner.Scan()

	// read an string from scanner
	op := scanner.Text()

	// operation
	operation := ""

	// split the string
	values := strings.Split(op, operation)

	// convert the values to int
	a, erra := strconv.Atoi(values[0]) // Atoi returns an int and an error
	b, errb := strconv.Atoi(values[1]) // we can scape the error with _

	// manipulating the errors
	if (erra != nil) || (errb != nil) {
		fmt.Println("Error: invalid input")
		os.Exit(1)
	}

	// evaluating the operation and printing the result
	switch operation {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		fmt.Println(a / b)
	default:
		fmt.Println("Error: invalid operation")
	}
}
