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
				if overlap([]byte(seq[4]), []byte(vl)) && overlap([]byte(seq[1]), []byte(vl)) {
					seq[9] = vl
				}
				if !overlap([]byte(seq[1]), []byte(vl)) {
					seq[6] = vl
				}
				if overlap([]byte(seq[1]), []byte(vl)) && !overlap([]byte(seq[4]), []byte(vl)) {
					seq[0] = vl
				}
			}
		}

		fmt.Println("break")
	}
}

func overlap(x, y []byte) bool {
	ret := false
	for _, val := range x {
		found := false
		for _, vl := range y {
			if vl == val {
				found = true
			}
		}
		if found {
			ret = true
		}
	}

	return ret
}
