package main

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

	// print the string
	fmt.Println(op)

	// split the string
	values := strings.Split(op, "+")

	// print the values
	fmt.Println(values)

	// print the sum (but isn't a sum, it's a concatenation)
	fmt.Println(values[0] + values[1])

	// convert the values to int
	a, _ := strconv.Atoi(values[0]) // Atoi returns an int and an error
	b, _ := strconv.Atoi(values[1]) // we can scape the error with _

	// print the sum
	fmt.Println(a + b)
}
