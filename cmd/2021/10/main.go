package main

import (
	"bufio"
	"fmt"
	"os"
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

	pairs['}'] = '{'
	pairs['>'] = '<'
	pairs[')'] = '('
	pairs[']'] = '['

	points[')'] = 3
	points[']'] = 57
	points['}'] = 1197
	points['>'] = 25137

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
	}

	fmt.Printf("Part 1: %d\n", total)

}
