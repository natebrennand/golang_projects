package main

import (
	"os"
	"fmt"
	"strconv"
)

func fib (arr []uint64, i int) []uint64 {
	// return if done
	if i == len(arr) {
		return arr
	}

	arr[i] = arr[i-1] + arr[i-2]
	return fib(arr, i+1)
}

func usuage() {
	fmt.Printf("usuage: go run fibonacci.go <int>\n")
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

	results := make([]uint64, n)
	results[0] = 0
	results[1] = 1

	results = fib(results, 2)
	fmt.Println(results)
}
