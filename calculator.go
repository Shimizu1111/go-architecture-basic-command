package main

import (
	"fmt"
	"os"
	"strconv"
)

func calculator(args []string) {
	if len(args) < 4 {
		fmt.Println("Usage: go run . [add|sub] number1 number2")
		os.Exit(1)
	}

	operation := args[1]
	num1, err1 := strconv.Atoi(args[2])
	num2, err2 := strconv.Atoi(args[3])

	if err1 != nil || err2 != nil {
		fmt.Println("Please provide valid numbers.")
		os.Exit(1)
	}

	switch operation {
	case "add":
		result := add(num1, num2)
		fmt.Printf("Result: %d\n", result)
	case "sub":
		result := subtract(num1, num2)
		fmt.Printf("Result: %d\n", result)
	default:
		fmt.Println("Invalid operation. Please use 'add' or 'sub'.")
		os.Exit(1)
	}
}
