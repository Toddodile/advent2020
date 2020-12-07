package dayOne

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Problem1() {
	fmt.Println("Path to expense report: ")

	var reportPath string

	fmt.Scan(&reportPath)
	if reportPath == "" {
		fmt.Println("Report path cannot be empty")
		os.Exit(1)
	}

	if _, err := os.Stat(reportPath); err != nil {
		fmt.Println("Report file specified does not exist")
		os.Exit(2)
	}

	file, err := os.Open(reportPath)
	if err != nil {
		fmt.Println("Error opening report file")
		os.Exit(3)
	}
	values, err := readInts(file)

	fmt.Println(productOfSums(values, 2020))
}

func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func productOfSums(values []int, sumTarget int) string {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i]+values[j] == sumTarget {
				return strconv.Itoa(values[i] * values[j])
			}
		}
	}
	return "Sum not found"
}
