package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	nums = map[string]bool{
		"0": true,
		"1": true,
		"2": true,
		"3": true,
		"4": true,
		"5": true,
		"6": true,
		"7": true,
		"8": true,
		"9": true,
	}
)

func main() {
	grid := parse()

	currstreak := ""
	streakvalid := false
	instreak := false

	validnums := []string{}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			itm := grid[y][x]
			_, found := nums[itm]
			if found {
				currstreak += itm
				instreak = true
				valid := checkNeighbours(x, y, grid)
				if valid {
					streakvalid = true
				}

				continue
			}
			if streakvalid {
				validnums = append(validnums, currstreak)
				streakvalid = false
			}
			if instreak {
				instreak = false
				fmt.Println(currstreak)
				currstreak = ""
			}

		}

	}
	fmt.Printf("Part 1: %v\n", parseValidNumbers(validnums))
}

func parseValidNumbers(nums []string) int {
	total := 0
	for _, c := range nums {
		n, err := strconv.Atoi(c)
		if err != nil {
			panic("Unable to parse number found")
		}
		total += n
	}

	return total
}

func checkNeighbours(x int, y int, grid [][]string) bool {
	neighbours := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for _, pairs := range neighbours {
		xp := x + pairs[0]
		yp := y + pairs[1]

		// Out of bound check
		if yp < 0 || yp >= len(grid) || xp < 0 || xp >= len(grid[y]) {
			continue
		}

		char := grid[yp][xp]

		// Check for number
		_, found := nums[char]
		if found {
			continue
		}

		if char != "." {
			return true
		}

	}
	return false
}

func parse() [][]string {
	var grid [][]string
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var line []string
		for _, c := range scanner.Text() {
			line = append(line, string(c))
		}
		grid = append(grid, line)
	}
	return grid
}
