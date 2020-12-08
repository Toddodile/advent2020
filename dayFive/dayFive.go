package dayFive

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type boardingPass struct {
	specifier string
	row       int
	column    int
}

func newBoardingPass(specifier string) boardingPass {
	return boardingPass{specifier: specifier}
}

func (bp *boardingPass) binarySearch(start int, stop int, path string) int {
	if start == stop {
		return start
	}
	if path[0] == 'F' || path[0] == 'L' {
		return bp.binarySearch(start, start+(stop-start)/2, path[1:])
	}
	return bp.binarySearch((start+(stop-start)/2)+1, stop, path[1:])
}

func (bp *boardingPass) getSeatId() int {
	return bp.row*8 + bp.column
}

func ScanPasses() {
	boardingPassFilePath := "boardingpass.txt"

	file, err := os.Open(boardingPassFilePath)
	if err != nil {
		fmt.Println("Error opening boarding pass file")
		os.Exit(3)
	}
	boardingPasses := parseBoardingPassFile(file)
	biggestId := 0
	seats := make(map[int]int)
	for _, bp := range boardingPasses {
		seats[bp.getSeatId()] = 1
		if bp.getSeatId() > biggestId {
			biggestId = bp.getSeatId()
		}
	}
	for i := 1; i < biggestId-1; i++ {
		if seats[i-1] == 1 && seats[i] == 0 && seats[i+1] == 1 {
			fmt.Println("My seat is " + strconv.Itoa(i))
			break
		}
	}
	fmt.Println("Largest Seat ID: " + strconv.Itoa(biggestId))
}

func parseBoardingPassFile(reader io.Reader) []boardingPass {
	var boardingPasses []boardingPass

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		bp := newBoardingPass(scanner.Text())
		bp.row = bp.binarySearch(0, 127, bp.specifier[:7])
		bp.column = bp.binarySearch(0, 7, bp.specifier[7:])
		boardingPasses = append(boardingPasses, bp)
	}
	return boardingPasses
}
