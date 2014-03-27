package main

import (
	"fmt"
	"os"
	"strconv"
)

type Unit struct {
	name string
	diff int    // in cents
	count int
}

func (u Unit) give (change int) bool {
	if u.diff <= change {
		return true
	}
	return false
}

func (u Unit) toString () string {
	if u.count > 0 {
		return fmt.Sprintf("%s: %d ", u.name, u.count)
	}
	return ""
}

func GetChange(tendered, cost int) []Unit {
	changeOptions := []Unit{
		Unit{"twenty", 2000, 0},
		Unit{"ten", 1000, 0},
		Unit{"five", 500, 0},
		Unit{"one", 100, 0},
		Unit{"quarter", 25, 0},
		Unit{"dime", 10, 0},
		Unit{"nickel", 5, 0},
		Unit{"penny", 1, 0},
	}
	diff := tendered - cost

	for {
		for i, amount := range(changeOptions) {
			if amount.give(diff) {
				changeOptions[i].count += 1
				diff -= amount.diff
				break
			}
		}
		if diff == 0 {
			break
		}
	}
	return changeOptions
}

func Usuage() {
	fmt.Println("USUAGE: go run change.go <AMOUNT TENDERED> <PRICE>")
	os.Exit(1)
}

func CliArgs() (int, int) {
	if len(os.Args) != 3 {
		Usuage()
	}
	tendered_float, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println(err.Error())
		Usuage()
	}
	tendered := int(100 * tendered_float)

	cost_float, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println(err.Error())
		Usuage()
	}
	cost := int(100 * cost_float)

	if tendered < cost {
		fmt.Println("You need to pay more than the item costs!")
		Usuage()
	}

	return tendered, cost
}

func main() {
	tendered, cost := CliArgs()
	changeGiven := GetChange(tendered, cost)

	outcome := ""
	for _, amount := range(changeGiven) {
		outcome += amount.toString()
	}
	fmt.Println(outcome)
}
