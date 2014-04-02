package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
)

const MAX_POWER = 7

type Result struct {
	num, count int
}

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

func maxCollatzRangeWorker(s int, max chan Result) {
	var n int
	magnitude := int(math.Pow10(MAX_POWER - 1))
	maxC := 0
	for i := s * magnitude; i < (s+1)*magnitude; i++ {
		temp := collatz(i, 0)
		if temp > maxC {
			n = i
			maxC = temp
		}
	}
	max <- Result{n, maxC}
}

func maxCollatz() {
	var maxNum, maxCount int
	maxs := make(chan Result)
	for i := 1; i < 10; i++ {
		go maxCollatzRangeWorker(i, maxs)
	}

	for i := 1; i < 10; i++ {
		r := <-maxs
		if r.count > maxCount {
			maxCount = r.count
			maxNum = r.num
		}
	}

	fmt.Printf("%d is the number with the most collatz iterations between 1 and %d with %d\n", maxNum, int(math.Pow10(MAX_POWER)), maxCount)
}

func main() {
	if len(os.Args) != 2 {
		runtime.GOMAXPROCS(11)
		maxCollatz()
	} else {
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("You must pass in an integer as an argument!\n")
			os.Exit(1)
		}

		fmt.Printf("The Collatz number of %d is %d.\n", n, collatz(n, 0))
	}
}
