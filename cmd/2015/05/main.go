package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	var nicestrings int

	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		isnice := isNice(line)

		if isnice {
			nicestrings++
		}

	}

	fmt.Printf("Nice strings: %d\n", nicestrings)
}
func isNice(line string) bool {
	if countVowels(line) && countLettersInRow(line) && !containsForbiddenString(line) {
		return true
	}
	return false
}

func countVowels(line string) bool {
	vowels := regexp.MustCompile("a|e|o|u|i")
	matches := vowels.FindAllStringIndex(line, -1)
	if len(matches) < 3 {
		return false
	}

	return true
}

func countLettersInRow(line string) bool {
	var prev rune
	for _, let := range line {
		if let == prev {
			return true
		}
		prev = let
	}

	return false
}

func containsForbiddenString(line string) bool {
	forbidden := regexp.MustCompile("ab|cd|pq|xy")
	matches := forbidden.FindAllStringIndex(line, -1)
	if len(matches) != 0 {
		return true
	}

	return false
}
