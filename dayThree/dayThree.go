package dayThree

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Navigate() {
	treeMapPath := "treemap.txt"

	file, err := os.Open(treeMapPath)
	if err != nil {
		fmt.Println("Error opening tree map file")
		os.Exit(3)
	}
	treeLines := parseTreeLines(file)
	oneByOne := findTrees(treeLines, 1, 1)
	threeByOne := findTrees(treeLines, 3, 1)
	fiveByOne := findTrees(treeLines, 5, 1)
	sevenByOne := findTrees(treeLines, 7, 1)
	oneByTwo := findTrees(treeLines, 1, 2)
	productOfTrees := oneByOne * threeByOne * fiveByOne * sevenByOne * oneByTwo
	fmt.Println("Product of trees: " + strconv.Itoa(productOfTrees))
}

func parseTreeLines(reader io.Reader) []string {
	var treeLines []string
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		treeLines = append(treeLines, scanner.Text())
	}
	return treeLines
}

func findTrees(treeLines []string, over int, down int) int {
	treeCount := 0
	offset := over
	for i := down; i < len(treeLines); i = i + down {
		if treeLines[i][offset] == '#' {
			treeCount++
		}
		offset = (offset + over) % len(treeLines[i])
	}
	fmt.Println("Trees over " + strconv.Itoa(over) + " down " + strconv.Itoa(down) + ": " + strconv.Itoa(treeCount))
	return treeCount
}
