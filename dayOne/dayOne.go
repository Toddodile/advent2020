package dayOne

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Problem1() {

	reportPath := "report.txt"

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
	for i := 0; i < len(values)-2; i++ {
		for j := i + 1; j < len(values)-1; j++ {
			for k := j + 1; k < len(values); k++ {
				if values[i]+values[j]+values[k] == sumTarget {
					return strconv.Itoa(values[i] * values[j] * values[k])
				}
			}
		}
	}
	return "Sum not found"
}
