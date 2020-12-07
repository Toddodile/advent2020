package dayTwo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func CorrectPasswords() {
	passwordPath := "password.txt"

	file, err := os.Open(passwordPath)
	if err != nil {
		fmt.Println("Error opening password file")
		os.Exit(3)
	}

	occurances, letters, passwords := parsePasswordFile(file)
	evaluatePasswordsSled(occurances, letters, passwords)
	evaluatePasswordsToboggan(occurances, letters, passwords)
}

func parsePasswordFile(reader io.Reader) ([][]int, []byte, []string) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	var occurances [][]int
	var letters []byte
	var passwords []string
	i := 0
	for scanner.Scan() {
		switch i {
		case 0:
			occurances = append(occurances, parseOccurance(scanner.Text()))
			i++
		case 1:
			letters = append(letters, scanner.Text()[0])
			i++
		case 2:
			passwords = append(passwords, scanner.Text())
			i = 0
		default:
			// Should never get here
			os.Exit(4)
		}
	}
	return occurances, letters, passwords
}

func parseOccurance(text string) []int {
	var occurance []int
	minmax := strings.Split(text, "-")
	if len(minmax) != 2 {
		min := 0
		max := 0
		occurance = append(occurance, min)
		occurance = append(occurance, max)
	} else {
		min, err := strconv.Atoi(minmax[0])
		max, err := strconv.Atoi(minmax[1])
		if err != nil {
			os.Exit(5)
		}
		occurance = append(occurance, min)
		occurance = append(occurance, max)
	}
	return occurance
}

func evaluatePasswordsSled(occurances [][]int, letters []byte, passwords []string) {
	if len(occurances) != len(letters) || len(letters) != len(passwords) {
		fmt.Println("Mismatching length of arrays")
		os.Exit(2)
	}

	count := 0
	for i := 0; i < len(occurances); i++ {
		min := occurances[i][0]
		max := occurances[i][1]
		letterCount := 0
		for j := 0; j < len(passwords[i]); j++ {
			if passwords[i][j] == letters[i] {
				letterCount++
			}
		}
		if letterCount >= min && letterCount <= max {
			count++
		}
	}
	fmt.Println("Sled Co valid password count: " + strconv.Itoa(count))
}

func evaluatePasswordsToboggan(indexes [][]int, letters []byte, passwords []string) {
	if len(indexes) != len(letters) || len(letters) != len(passwords) {
		fmt.Println("Mismatching length of arrays")
		os.Exit(2)
	}

	count := 0
	for i := 0; i < len(indexes); i++ {
		pos1 := indexes[i][0] - 1
		pos2 := indexes[i][1] - 1
		if (passwords[i][pos1] == letters[i] || (pos2 < len(passwords[i]) && passwords[i][pos2] == letters[i])) && passwords[i][pos1] != passwords[i][pos2] {
			count++
		}
	}
	fmt.Println("Toboggan Co valid password count: " + strconv.Itoa(count))
}
