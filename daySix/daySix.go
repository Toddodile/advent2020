package daySix

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func TrueFalseCounting() {
	customsFilePath := "customs.txt"

	file, err := os.Open(customsFilePath)
	if err != nil {
		fmt.Println("Failed to open customs file")
		os.Exit(3)
	}
	answerSheetsOneof, answerSheetsAllOf := parseCustomsFile(file)
	sumOfTruesOneOf := 0
	sumOfTruesAllOf := 0
	//alphabet := "abcdefghijklmnopqrstuvwxyz"
	for i, answers := range answerSheetsOneof {
		sumOfTruesOneOf += len(answers)
		for _, value := range answerSheetsAllOf[i] {
			sumOfTruesAllOf += value
		}
	}
	fmt.Println("Sum of True-For-One values: " + strconv.Itoa(sumOfTruesOneOf))
	fmt.Println("Sum of True-For-All values: " + strconv.Itoa(sumOfTruesAllOf))
}

func parseCustomsFile(reader io.Reader) ([]map[rune]int, []map[rune]int) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	var answerSheetsOneOf []map[rune]int
	var answerSheetsAllOf []map[rune]int
	answersOneOf := make(map[rune]int)
	answersAllOf := make(map[rune]int)
	groupSize := 0
	for scanner.Scan() {
		line := scanner.Text()
		//Separator
		if line == "" {
			answerSheetsOneOf = append(answerSheetsOneOf, answersOneOf)
			for i := 'a'; i <= 'z'; i++ {
				if answersAllOf[i] == groupSize {
					answersAllOf[i] = 1
				} else {
					answersAllOf[i] = 0
				}
			}
			answerSheetsAllOf = append(answerSheetsAllOf, answersAllOf)
			groupSize = 0
			answersOneOf = make(map[rune]int)
			answersAllOf = make(map[rune]int)
		} else {
			groupSize++
			for _, char := range line {
				answersOneOf[char] = 1
				answersAllOf[char]++
			}
		}
	}
	answerSheetsOneOf = append(answerSheetsOneOf, answersOneOf)
	for i := 'a'; i <= 'z'; i++ {
		if answersAllOf[i] == groupSize {
			answersAllOf[i] = 1
		} else {
			answersAllOf[i] = 0
		}
	}
	answerSheetsAllOf = append(answerSheetsAllOf, answersAllOf)
	return answerSheetsOneOf, answerSheetsAllOf
}
