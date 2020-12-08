package main

import (
	"fmt"

	dayFive "github.com/todd5913/advent2020/dayFive"
	dayFour "github.com/todd5913/advent2020/dayFour"
	dayOne "github.com/todd5913/advent2020/dayOne"
	daySix "github.com/todd5913/advent2020/daySix"
	dayThree "github.com/todd5913/advent2020/dayThree"
	dayTwo "github.com/todd5913/advent2020/dayTwo"
)

func main() {
	fmt.Println("Pick a day (1, 2, etc): ")

	var dayChoice string
	//fmt.Scan(&dayChoice)
	dayChoice = "6"

	switch dayChoice {
	case "1":
		dayOne.Problem1()
	case "2":
		dayTwo.CorrectPasswords()
	case "3":
		dayThree.Navigate()
	case "4":
		dayFour.ScanPasswords()
	case "5":
		dayFive.ScanPasses()
	case "6":
		daySix.TrueFalseCounting()
	default:
		fmt.Println("That day doesn't exist yet")
	}
}
