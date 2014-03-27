package main

import (
	"fmt"
	"os"
	"strconv"
	"math"
)

const (
	HappyAttempts = 100
	NumFind = 8
)

func process (n int, total *int) int {
	if n == 0 {
		return *total
	}
	*total += int(math.Pow(float64(n % 10), float64(2)))
	return process(n/10, total)
}

func isHappy (n int) bool {
	var iterate int
	for iterate = HappyAttempts; iterate > 0; iterate-- {
		total := 0
		n = process(n, &total)
		if n == 1 {
			return true
		}
	}
	return false
}

// Prints the proper usuage of the program before exiting on error
func Usuage() {
	fmt.Println("USUAGE: go run happy.go <N>\n")
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
	sol := 0

	fmt.Printf("Started with %d, first %d \"happy\" numbers are:\n", n, NumFind)
	for ; sol < NumFind; {
		ok := isHappy(n)
		if ok {
			fmt.Printf("%d\n", n)
			sol++
		}
		n++
	}
}
