package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	total := 0
	var last []rune
	pairs := make(map[rune]rune)
	points := make(map[rune]int)
	p2points := make(map[rune]int)

	pairs['}'] = '{'
	pairs['>'] = '<'
	pairs[')'] = '('
	pairs[']'] = '['

	points[')'] = 3
	points[']'] = 57
	points['}'] = 1197
	points['>'] = 25137

	p2points['('] = 1
	p2points['['] = 2
	p2points['{'] = 3
	p2points['<'] = 4

	var correctlines []string
	for scanner.Scan() {
		var syntaxerror bool
		for _, char := range scanner.Text() {
			if (char == '}' || char == ')' || char == '>' || char == ']') && last[len(last)-1] != pairs[char] {
				// Incorrect closing char found
				total += points[char]
				syntaxerror = true
			}
			if (char == '}' || char == ')' || char == '>' || char == ']') && last[len(last)-1] == pairs[char] {
				// Incorrect closing char found
				last = last[:len(last)-1]
			}
			if char == '{' || char == '(' || char == '<' || char == '[' {
				last = append(last, char)
			}
			if syntaxerror {
				break
			}
		}
		if !syntaxerror {
			correctlines = append(correctlines, scanner.Text())
		}

	}

	fmt.Printf("Part 1: %d\n", total)

	var incpoints []int
	for _, line := range correctlines {
		var score int
		var hist []rune
		for _, char := range line {
			if (char == '}' || char == ')' || char == '>' || char == ']') && hist[len(hist)-1] == pairs[char] {
				// Closing char found
				hist = hist[:len(hist)-1]
			}
			if char == '{' || char == '(' || char == '<' || char == '[' {
				hist = append(hist, char)
			}
		}

		for i := len(hist) - 1; i >= 0; i-- {
			if hist[i] == '{' || hist[i] == '(' || hist[i] == '<' || hist[i] == '[' {
				fmt.Printf("%v", string(hist[i]))
				score = score * 5
				score += p2points[hist[i]]
			}
		}

		incpoints = append(incpoints, score)
		fmt.Printf(" - Score: %d\n", score)
	}

	sort.Ints(incpoints)

	fmt.Printf("Part 2: %d\n", incpoints[len(incpoints)/2])
	fmt.Printf("bla: %d\n", len(incpoints))
	fmt.Printf("Part 2: %d\n", incpoints)

}
