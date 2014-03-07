package main

import (
	"fmt"
	"os"
	"strconv"
)

// Calculates the cost and prints the calculations
func FindCost(w, h, c int) int {
	fmt.Printf("%d X %d @ %d/sq ft = ", w, h, c)
	cost := w * h * c
	fmt.Printf("%d\n", cost)
	return cost
}

// Prints the proper usuage of the program before exiting on error
func Usuage() {
	fmt.Println("USUAGE: go run floor.go <WIDTH> <HEIGHT> <COST / sq foot>")
	os.Exit(1)
}

// converts a specified CLI arg to an integer
func ArgsToInt(index int) int {
	x, err := strconv.Atoi(os.Args[index])
	if err != nil {
		Usuage()
	}
	return x
}

func main() {
	if len(os.Args) != 4 {
		Usuage()
	}

	width, height, cost := ArgsToInt(1), ArgsToInt(2), ArgsToInt(3)
	FindCost(width, height, cost)
}
