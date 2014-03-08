package main

import (
	"fmt"
	"os"
	"strconv"
)

func GetChange (tendered, cost int) map[string]int {
	change := make(map[string]int)

	diff := tendered - cost
	for {
		switch {
		case diff >= 2000:
			_, ok := change["twenty"]
			if !ok {
				change["twenty"] = 0
			}
			change["twenty"]++
			diff  -= 2000
		case diff >= 1000:
			_, ok := change["ten"]
			if !ok {
				change["ten"] = 0
			}
			change["ten"]++
			diff  -= 1000
		case diff >= 500:
			_, ok := change["five"]
			if !ok {
				change["five"] = 0
			}
			change["five"]++
			diff  -= 500
		case diff >= 100:
			_, ok := change["one"]
			if !ok {
				change["one"] = 0
			}
			change["one"]++
			diff  -= 100
		case diff >= 25:
			_, ok := change["quarter"]
			if !ok {
				change["quarter"] = 0
			}
			change["quarter"]++
			diff  -= 25
		case diff >= 10:
			_, ok := change["dime"]
			if !ok {
				change["dime"] = 0
			}
			change["dime"]++
			diff  -= 10
		case diff >= 5:
			_, ok := change["nickel"]
			if !ok {
				change["nickel"] = 0
			}
			change["nickel"]++
			diff  -= 5
		case diff >= 1:
			_, ok := change["penny"]
			if !ok {
				change["penny"] = 0
			}
			change["penny"]++
			diff  -= 1
		default:
			return change
		}
	}
	return change
}

func Usuage() {
	fmt.Println("USUAGE: go run change.go <AMOUNT TENDERED> <PRICE>")
	os.Exit(1)
}

func CliArgs() (int, int){
	if len(os.Args) != 3 {
		Usuage()
	}
	tendered_float, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println(err.Error())
		Usuage()
	}
	tendered := int(100*tendered_float)

	cost_float, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println(err.Error())
		Usuage()
	}
	cost := int(100*cost_float)


	if tendered < cost {
		fmt.Println("You need to pay more than the item costs!")
		Usuage()
	}

	return tendered, cost
}

func main() {
	tendered, cost := CliArgs()
	fmt.Println(GetChange(tendered, cost))
}
