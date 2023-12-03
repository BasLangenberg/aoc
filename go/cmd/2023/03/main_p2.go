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

	var gears [][]int

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == "*" {
				gear := []int{y, x}
				gears = append(gears, gear)
			}

		}
	}
	fmt.Printf("Part 2: %v\n", multiplyNeighbours(grid, gears))
	fmt.Println("Too low: 76255482")
}

func multiplyNeighbours(grid [][]string, coords [][]int) int {
	total := 0
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

	for _, coord := range coords {
		products := make(map[string]bool)
		for _, pairs := range neighbours {

			xp := coord[1] + pairs[0]
			yp := coord[0] + pairs[1]

			// Out of bound check
			if yp < 0 || yp >= len(grid) || xp < 0 || xp >= len(grid[coord[0]]) {
				continue
			}

			char := grid[yp][xp]

			// Check for number
			_, found := nums[char]
			if found {
				step := 0
				for {
					step++
					if xp-step < 0 {
						break
					}
					lchar := grid[yp][xp-step]
					_, err := strconv.Atoi(lchar)
					if err != nil {
						break
					}
				}

				numbr := ""
				for {
					step--
					if xp-step > len(grid[0]) -1 {
						break
					}
					lchar := grid[yp][xp-step]
					_, err := strconv.Atoi(lchar)
					if err != nil {
						step--
						break
					}
					numbr += lchar
				}
				products[numbr] = true
			}

		}
		var mult []int
		for k := range products {
			m, _ := strconv.Atoi(k)
			mult = append(mult, m)
		}
		fmt.Println(mult)
		if len(mult) < 2 {
			continue
		}

		ratio := 1
		for _, num := range mult {
			ratio = ratio * num
		}

		total += ratio
	}
	return total
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
