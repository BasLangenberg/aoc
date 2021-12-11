package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type octo struct {
	level, flashes   int
	flashed, counted bool
}

func (o *octo) inc() {
	o.level++
	if o.level > 9 && !o.flashed {
		o.flashed = true
		o.flashes++
	}
}

func (o *octo) reset() {
	if o.flashed {
		o.level = 0
		o.flashed = false
		o.counted = false
	}
}

func checkdiag(octi [][]octo) bool {
	updates := false
	for key, octos := range octi {
		for k, octo := range octos {
			if octo.flashed && !octo.counted {
				updates = true
				octi[key][k].counted = true
				// left
				if k-1 >= 0 {
					octi[key][k-1].inc()
				}
				// left up
				if k-1 >= 0 && key-1 >= 0 {
					octi[key-1][k-1].inc()
				}
				// up
				if key-1 >= 0 {
					octi[key-1][k].inc()
				}
				// right up
				if k+1 <= 9 && key-1 >= 0 {
					octi[key-1][k+1].inc()
				}
				// right
				if k+1 <= 9 {
					octi[key][k+1].inc()
				}
				// right down
				if k+1 <= 9 && key+1 <= 9 {
					octi[key+1][k+1].inc()
				}
				// down
				if key+1 <= 9 {
					octi[key+1][k].inc()
				}
				// left down
				if k-1 >= 0 && key+1 <= 9 {
					octi[key+1][k-1].inc()
				}
			}
		}
	}
	return updates
}

func main() {
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var grid [][]octo
	for scanner.Scan() {
		var line []octo
		for _, char := range scanner.Text() {
			i, _ := strconv.Atoi(string(char))
			o := octo{
				level: i,
			}
			line = append(line, o)

		}

		grid = append(grid, line)
	}

	for i := 0; i < 100; i++ {
		for key, val := range grid {
			for k := range val {
				grid[key][k].inc()
			}
		}
		for {
			if !checkdiag(grid) {
				break
			}
		}

		for key, val := range grid {
			for k := range val {
				grid[key][k].reset()
			}
		}

		// Print
		for key, val := range grid {
			for k := range val {
				fmt.Printf("%d", grid[key][k].level)
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")

	}

	var flashcount int
	for _, val := range grid {
		for _, v := range val {
			flashcount += v.flashes
		}
	}

	fmt.Printf("Part 1: %d flashed\n", flashcount)

}
