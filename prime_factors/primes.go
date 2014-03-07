package main

import (
	"fmt"
	"os"
	"strconv"
)

func FindPrimes(results []int, n int) []int {
	// finds a prime factors of a number
	i := 2
	for {
		if i > n {
			return results
		}
		if n%i == 0 {
			n /= i
			// skip repeats
			if len(results) == 0 || results[len(results)-1] != i {
				results = append(results, i)
			}
		} else {
			i++
		}
	}
}

func usuage() {
	fmt.Printf("usuage: go run primes.go <int>\n")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		usuage()
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usuage()
	}

	subPrimes := make([]int, 0, 25)
	fmt.Println(FindPrimes(subPrimes, n))
	os.Exit(0)
}
