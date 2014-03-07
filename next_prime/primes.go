package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// IsPrime determines if an integer is a prime number
func IsPrime(n int) bool {
	limit := int(math.Sqrt(float64(n)))
	if n == 1 || n == 2 { // base cases
		return true
	}

	i := 3
	for {
		if i > limit {
			break
		}
		if n%i == 0 {
			return false
		}
		i += 2
	}
	return true
}

// NextPrime repeatedly calls IsPrime() to find the next prime number
func NextPrime(n int) int {
	if n%2 == 0 { // make n odd
		n--
	}

	for {
		n += 2
		if IsPrime(n) {
			return n
		}
	}
}

// Usuage prints the proper way to call the function before exiting with an error
func Usuage() {
	fmt.Println("USUAGE: go run primes.go <int>")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		Usuage()
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		Usuage()
	}

	fmt.Println(NextPrime(n))
	os.Exit(0)
}
