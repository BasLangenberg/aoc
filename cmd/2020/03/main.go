package main

import (
	"bufio"
	"fmt"
	"os"
)

var gamemap []string

func main() {
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	// load map
	for scanner.Scan() {
		gamemap = append(gamemap, scanner.Text())
	}

	for i := 0; i < 10; i++ {
		for num := range gamemap {
			gamemap[num] = gamemap[num] + gamemap[num]
		}
	}

	// Part A
	fmt.Printf("Part A: Amount of trees hit: %v\n", checkSlope([]int{1, 3}))

	// Part B

	//Right 1, down 1.
	//Right 3, down 1. (This is the slope you already checked.)
	//Right 5, down 1.
	//Right 7, down 1.
	//Right 1, down 2.

	slopes := [][]int{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	}

	var trees []int
	for _, slope := range slopes {
		trees = append(trees, checkSlope(slope))
	}

	amount := 1
	for i := 1; i < len(trees); i++ {
		amount = amount * trees[i]
	}

	fmt.Printf("Part B: Amount of trees hit: %v\n", amount)
}

func checkSlope(slope []int) int {
	coords := [][]int{[]int{0, 0}}

	for i := 0; i < len(gamemap)-1; i++ {
		currentCords := make([]int, 2)
		copy(currentCords, coords[len(coords)-1])

		if currentCords[0]+slope[0] >= len(gamemap)-1 {
			currentCords[0] = currentCords[0] + 1
			coords = append(coords, currentCords)
			break
		}
		currentCords[0] = currentCords[0] + slope[0]
		currentCords[1] = currentCords[1] + slope[1]

		coords = append(coords, currentCords)
	}

	var trees int

	for _, cord := range coords {
		if string(gamemap[cord[0]][cord[1]]) == "#" {
			trees++
		}
	}

	return trees
}
