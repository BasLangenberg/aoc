package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func main() {
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	coords := make(map[int][][]int)
	var num int

	for scanner.Scan() {
		crds := strings.Split(scanner.Text(), " -> ")
		for _, val := range crds {
			tpl := strings.Split(val, ",")
			one, _ := strconv.Atoi(tpl[0])
			two, _ := strconv.Atoi(tpl[1])
			coords[num] = append(coords[num], []int{one, two})
		}
		num++
	}

	gamemap := make(map[point]int)

	for _, val := range coords {
		first := point{x: val[0][0], y: val[0][1]}
		second := point{x: val[1][0], y: val[1][1]}

		// part 1
		if first.x == second.x || first.y == second.y {
			deltax, deltay := 0, 0

			if first.x < second.x {
				deltax = 1
			}

			if first.x > second.x {
				deltax = -1
			}

			if first.y < second.y {
				deltay = 1
			}

			if first.y > second.y {
				deltay = -1
			}

			init := first
			for {
				gamemap[init]++
				if init.x == second.x && init.y == second.y {
					break
				}

				init.x += deltax
				init.y += deltay
			}
		}

	}
	var counter int

	for _, val := range gamemap {
		if val > 1 {
			counter++
		}
	}
	fmt.Printf("Part 1: %d", counter)
}
