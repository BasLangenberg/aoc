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

	var unique int
	seq := make(map[string]int)

	for _, val := range lines {
		for _, vl := range strings.Split(val[1], " ") {
			if len(vl) == 2 || len(vl) == 3 || len(vl) == 4 || len(vl) == 7 {
				seq[vl]++
				unique++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", unique)

	for key, val := range seq {
		if len(key) == 2 {
			fmt.Printf("Key: %v, Val: %v\n", key, val)
		}
	}

}
