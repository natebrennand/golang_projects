package main

import (
	"fmt"
	"os"
	"strconv"
)

// Collatz Conjecture
// Start with a number n > 1.
// Find the number of steps it takes to reach one using the following process:
//		If n is even, divide it by 2.
// 		If n is odd, multiply it by 3 and add 1.

func collatz(n, count int) int {
	if n == 1 {
		return count
	} else if n%2 == 0 {
		return collatz(n/2, count+1)
	} else {
		return collatz(n*3+1, count+1)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You must pass in an integer as an argument!\n")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("You must pass in an integer as an argument!\n")
		os.Exit(1)
	}

	fmt.Printf("The Collatz number of %d is %d.\n", n, collatz(n, 0))
}
