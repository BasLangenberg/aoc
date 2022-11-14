package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var lines [][]string

	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), " | "))
	}

	total := 0
	for _, val := range lines {
		all := append(strings.Split(val[0], " "), strings.Split(val[1], " ")...)

		seq := make(map[int]string)

		// Find known numbers
		for _, vl := range all {
			if len(vl) == 2 {
				seq[1] = vl
			}

			if len(vl) == 3 {
				seq[7] = vl
			}

			if len(vl) == 4 {
				seq[4] = vl
			}

			if len(vl) == 7 {
				seq[8] = vl
			}
		}

		// parse len6
		for _, vl := range all {
			if len(vl) == 6 {
				if overlap([]byte(seq[4]), []byte(vl)) {
					seq[9] = vl
					continue
				}
				if overlap([]byte(seq[7]), []byte(vl)) {
					seq[0] = vl
					continue
				}
				if !overlap([]byte(seq[4]), []byte(vl)) {
					seq[6] = vl
				}
			}
		}

		// parse len5
		for _, vl := range all {
			if len(vl) == 5 {
				if overlap([]byte(seq[1]), []byte(vl)) {
					seq[3] = vl
					continue
				}
				if overlap([]byte(vl), []byte(seq[6])) {
					seq[5] = vl
					continue
				}
				seq[2] = vl
			}
		}

		// get outcome
		counter := 0
		counting := 0
		for _, vl := range strings.Split(val[1], " ") {
			for key, val := range seq {
				if len(val) == len(vl) {
					if overlap([]byte(vl), []byte(val)) {
						switch counting {
						case 0:
							counter += key * 1000
						case 1:
							counter += key * 100
						case 2:
							counter += key * 10
						case 3:
							counter += key
						}
						fmt.Printf("%d", key)
						counting++
					}
				}
			}
		}
		total += counter

	}

	fmt.Printf("Part 2: %d", total)
}

func overlap(x, y []byte) bool {
	var foundcount int
	for _, val := range x {
		for _, vl := range y {
			if vl == val {
				foundcount++
				continue
			}
		}
	}

	return len(x) == foundcount
}
