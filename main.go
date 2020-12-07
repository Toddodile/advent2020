package main

import (
	"fmt"

	dayOne "github.com/todd5913/advent2020/dayOne"
)

func main() {
	fmt.Println("Pick a day (1, 2, etc): ")

	var dayChoice string
	fmt.Scan(&dayChoice)

	switch dayChoice {
	case "1":
		dayOne.Problem1()
	default:
		fmt.Println("That day doesn't exist yet")
	}
}
