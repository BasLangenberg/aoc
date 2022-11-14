package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type point struct {
	x, y int
}

var (
	directions = []point{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
)

func main() {
	input, err := os.Open("input")
	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	defer input.Close()

	var card [][]int

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var line []int
		for _, num := range []byte(scanner.Text()) {
			x, _ := strconv.Atoi(string(num))
			line = append(line, x)
		}
		card = append(card, line)
	}

	var basins []point
	for y, line := range card {
		for x := range line {
			lowest := true
			for _, d := range directions {
				np := point{x: x + d.x, y: y + d.y}
				if np.x >= 0 && np.x < len(card[y]) && np.y >= 0 && np.y < len(card) {
					if card[np.y][np.x] <= card[y][x] {
						lowest = false
						break
					}
				}
			}
			if lowest {
				basins = append(basins, point{x, y})
			}
		}
	}

	risklevel := 0
	for _, val := range basins {
		risklevel += 1 + card[val.y][val.x]
	}

	fmt.Printf("Part 1: %d\n", risklevel)

	// Part 2
	// Huge help from https://github.com/Skarlso/aoc2021/blob/main/day09/part2/main.go
	// And https://www.redblobgames.com/pathfinding/a-star/introduction.html
	var basinsizes []int

	for _, val := range basins {
		basinsizes = append(basinsizes, calcfrontier(val, card))
	}

	sort.Ints(basinsizes)
	fmt.Printf("Part 2: %d\n", basinsizes[len(basinsizes)-1]*basinsizes[len(basinsizes)-2]*basinsizes[len(basinsizes)-3])
}

func calcfrontier(p point, card [][]int) int {
	start := p
	seen := map[point]struct{}{
		start: {},
	}
	queue := []point{start}
	var current point
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		for _, next := range neighbors(current, card) {
			if _, ok := seen[next]; !ok {
				queue = append(queue, next)
				seen[next] = struct{}{}
			}
		}
	}

	return len(seen)
}

func neighbors(p point, card [][]int) []point {
	var result []point

	for _, d := range directions {
		np := point{x: p.x + d.x, y: p.y + d.y}
		if np.x >= 0 && np.x < len(card[p.y]) && np.y >= 0 && np.y < len(card) {
			if card[np.y][np.x] > card[p.y][p.x] && card[np.y][np.x] != 9 {
				result = append(result, np)
			}
		}
	}

	return result
}
