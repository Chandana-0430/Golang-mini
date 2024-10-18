package main

import (
	"fmt"
	"math"
)

func calculate(a float64, b float64, operator string) float64 {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b != 0 {
			return a / b
		} else {
			fmt.Println("ERROR: Division by zero")
			return math.NaN()
		}
	case "**":
		return math.Pow(a, b)
	default:
		fmt.Println("Invalid operator")
		return math.NaN()
	}
}

func main() {
	var a, b float64
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scan(&a)

	fmt.Print("Enter operator (+, -, *, /, **): ")
	fmt.Scan(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	result := calculate(a, b, operator)
	if !math.IsNaN(result) {
		fmt.Printf("Result: %.2f\n", result)
	}
}
