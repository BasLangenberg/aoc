package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}

	var card [][]int

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var line []int
		for _, num := range []byte(scanner.Text()) {
			x, _ := strconv.Atoi(string(num))
			line = append(line, x)
			fmt.Printf("%d", x)

		}
		fmt.Printf("\n")
		card = append(card, line)
	}

	risklevel := 0
	var lowpoints [][]int
	for y, line := range card {
		for x, val := range line {
			// Up
			if y != 0 {
				if card[y-1][x] <= val {
					continue
				}
			}

			// Down
			if y < len(card)-1 {
				if card[y+1][x] <= val {
					continue
				}
			}

			// Left
			if x > 0 {
				if card[y][x-1] <= val {
					continue
				}
			}

			// Right
			if x < len(line)-1 {
				if card[y][x+1] <= val {
					continue
				}
			}

			risklevel += 1 + val
			lowpoints = append(lowpoints, []int{x, y})

		}
	}

	fmt.Printf("Part 1: %d\n", risklevel)
	fmt.Printf("lowpoints: %+v", lowpoints)
	var basinsize []int
	for it, val := range lowpoints {
		basincoords := make(map[string]int)
		y, x := val[0], val[1]
		basincoords[fmt.Sprintf("%d-%d", y, x)] = 1

		// Up
		count := 0
		for {
			count++

			if y-count < 0 {
				break
			}
			if card[y-count][x] == 9 {
				break
			}

			basincoords[fmt.Sprintf("%d-%d", y-1, x)] = 1

		}
		// Down
		count = 0
		for {
			count++

			if y+count >= len(card) {
				break
			}
			if card[y+count][x] == 9 {
				break
			}

			basincoords[fmt.Sprintf("%d-%d", y+1, x)] = 1

		}

		// Left
		count = 0
		for {
			count++

			if x-count < 0 {
				break
			}
			if card[y][x-count] == 9 {
				break
			}

			basincoords[fmt.Sprintf("%d-%d", y, x-1)] = 1

		}

		// Right
		count = 0
		for {
			count++

			if count+x >= len(card[y]) {
				break
			}
			if card[y][x+count] == 9 {
				break
			}

			basincoords[fmt.Sprintf("%d-%d", y, x+1)] = 1

		}

		basinsize = append(basinsize, len(basincoords))
		for key, val := range basincoords {
			fmt.Printf("iter: %d - %v %+v\n", it, key, val)
		}

	}

	fmt.Println(basinsize)
	fmt.Println(len(lowpoints))

}
