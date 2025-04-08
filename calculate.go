package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Knetic/govaluate"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a mathematical expression (e.g., '3 + 4 * 2 / (1 - 5)^2^3'): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Trim the input to remove any surrounding whitespace
	input = strings.TrimSpace(input)

	// Create a new expression
	expression, err := govaluate.NewEvaluableExpression(input)
	if err != nil {
		fmt.Println("Error parsing expression:", err)
		return
	}

	// Evaluate the expression
	result, err := expression.Evaluate(nil)
	if err != nil {
		fmt.Println("Error evaluating expression:", err)
		return
	}

	fmt.Printf("Result: %v\n", result)
}
