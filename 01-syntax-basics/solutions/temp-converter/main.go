package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <temp> <unit>")
		fmt.Println("Example: go run main.go 25 C")
		return
	}

	tempStr := os.Args[1]
	unit := os.Args[2]

	temp, err := strconv.ParseFloat(tempStr, 64)
	if err != nil {
		fmt.Println("Error: invalid temperature value")
		return
	}

	var result float64
	var resultUnit string

	if unit == "C" {
		result = (temp * 9 / 5) + 32
		resultUnit = "F"
	} else if unit == "F" {
		result = (temp - 32) * 5 / 9
		resultUnit = "C"
	} else {
		fmt.Println("Error: Unit must be 'C' or 'F'")
		return
	}

	fmt.Printf("%.2f%s = %.2f%s\n", temp, unit, result, resultUnit)
}
