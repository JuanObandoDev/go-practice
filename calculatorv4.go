package main

// calculator using structs and methods

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(input, op string) int {
	// split the string
	values := strings.Split(input, op)
	a := parser(values[0])
	b := parser(values[1])

	// evaluating the operation and printing the result
	switch op {
	case "+":
		return (a + b)
	case "-":
		return (a - b)
	case "*":
		return (a * b)
	case "/":
		return (a / b)
	default:
		fmt.Println("Error: invalid operation")
	}
	return 0
}

func parser(input string) int {
	// convert the values to int
	a, _ := strconv.Atoi(input) // Atoi returns an int and an error
	return a
}

func reader() string {
	// creating scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)

	// activate scanner
	scanner.Scan()

	// read an string from scanner
	return scanner.Text()
}

func main() {
	// read an string from scanner
	input := reader()

	// operation
	operation := reader()

	// creating a calc struct
	c := calc{}
	fmt.Println(c.operate(input, operation))
}
