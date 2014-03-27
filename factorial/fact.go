package main

import (
	"fmt"
	"os"
	"strconv"
)

func recursiveFactorial (n int) int {
	if n==0 {
		return 1
	}
	return n * recursiveFactorial(n-1)
}

func loopFactorial (n int) int {
	num := 1
	for ; n > 0; n-- {
		num *= n
	}
	return num
}

// Prints the proper usuage of the program before exiting on error
func Usuage() {
	fmt.Println("USUAGE: go run fact.go <N>\n")
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
	if len(os.Args) != 2 {
		Usuage()
	}
	n := ArgsToInt(1)

	fmt.Printf("recursive: %d\n", recursiveFactorial(n))
	fmt.Printf("loop: %d\n", loopFactorial(n))
}
