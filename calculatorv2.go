package main

// calculator using conditionals an manipulating errors

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

	// split the string
	values := strings.Split(op, "+")

	// convert the values to int
	a, erra := strconv.Atoi(values[0]) // Atoi returns an int and an error
	b, errb := strconv.Atoi(values[1]) // we can scape the error with _

	// manipulating the errors
	if (erra != nil) || (errb != nil) {
		fmt.Println("Error: invalid input")
	} else {
		fmt.Println(a + b)
	}
}
