package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1]
	number, err := strconv.Atoi(input)
	if err != nil {
		panic(fmt.Sprint("Invalid input ", input, " is not integer"))
	}
	primeFactor(number)
}

func primeFactor(number int) {
	currentFactor := 0
	for number%2 == 0 {
		number = number / 2
		if currentFactor != 2 {
			currentFactor = 2
			fmt.Print(currentFactor, " ")
		}
	}

	for i := 3; i*i < number; i += 2 {
		for number%i == 0 {
			if currentFactor != i {
				currentFactor = i
				fmt.Print(currentFactor, " ")
			}
			number /= i
		}
	}

	if number > 2 {
		fmt.Print(number)
	}
	fmt.Println()
}
