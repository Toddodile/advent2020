package main

import (
	"fmt"

	dayOne "github.com/todd5913/advent2020/dayOne"
	dayThree "github.com/todd5913/advent2020/dayThree"
	dayTwo "github.com/todd5913/advent2020/dayTwo"
)

func main() {
	fmt.Println("Pick a day (1, 2, etc): ")

	var dayChoice string
	fmt.Scan(&dayChoice)

	switch dayChoice {
	case "1":
		dayOne.Problem1()
	case "2":
		dayTwo.CorrectPasswords()
	case "3":
		dayThree.Navigate()
	default:
		fmt.Println("That day doesn't exist yet")
	}
}
