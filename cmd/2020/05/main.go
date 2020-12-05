package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	memo := [][]int{}

	for scanner.Scan() {
		row, column, seatid := checkSeat(scanner.Text())
		memo = append(memo, []int{row, column, seatid})
	}

	var biggestSeatID int
	for _, bp := range memo {
		if bp[2] > biggestSeatID {
			biggestSeatID = bp[2]
		}
	}

	fmt.Printf("highest Seat ID: %v\n", biggestSeatID)

	// Part B
	// This part is super naive. I figured it out by looking at the output below
	// Sequences should be ignored! ;-)
	// Will probably not fix this later

	min := memo[0][2]

	for _, bp := range memo {
		if bp[2] < min {
			min = bp[2]
		}
	}

	for i := min; i < 839; i++ {
		var found bool
		for _, bp := range memo {
			if bp[2] == i {
				found = true
			}
		}
		if !found {
			fmt.Printf("%v is not on the list\n", i)
		}
	}
}

func checkSeat(input string) (int, int, int) {
	upRow := 127
	lowRow := 0
	upColumn := 7
	lowColumn := 0

	//	seatid = (row * 8) + column
	for i := 0; i < 7; i++ {
		mid := (upRow + lowRow) / 2
		switch input[i] {
		case 'F':
			upRow = mid
		case 'B':
			lowRow = mid
		default:
			fmt.Printf("Unable to parse %v as either F or B", input[i])
			os.Exit(1)
		}
	}

	for i := 7; i < 10; i++ {
		mid := (upColumn + lowColumn) / 2
		switch input[i] {
		case 'R':
			lowColumn = mid
		case 'L':
			upColumn = mid
		default:
			fmt.Printf("Unable to parse %v as either F or B", input[i])
			os.Exit(1)
		}
	}
	return upRow, upColumn, (upRow * 8) + upColumn
}
