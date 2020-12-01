package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const lowercasealphabet = "abcdefghijklmnopqrstuvwxyz"

func main() {
	var nicestrings int
	var nice2pointO int

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
		isnice2 := isNiceTwoPointZero(line)

		if isnice {
			nicestrings++
		}

		if isnice2 {
			nice2pointO++
		}

	}

	fmt.Printf("Nice strings: %d\n", nicestrings)
	fmt.Printf("Nice 2.0 strings: %d\n", nice2pointO)

}

func isNiceTwoPointZero(line string) bool {
	// We need two letter pairs, but they cannot overlap
	// e.g. eeyee is ok
	// e.g. yaaay is not ok!

	appearstwice := false
	oneletterinbetween := false

	pairs := make(map[string][]int)
	letters := make(map[string][]int)

	for i := 0; i < len(line)-1; i++ {
		pair := fmt.Sprintf(string(line[i]) + string(line[i+1]))
		pairs[pair] = append(pairs[pair], i)
	}

	for i := 0; i < len(line); i++ {
		letters[string(line[i])] = append(letters[string(line[i])], i)
	}

	for _, positions := range pairs {
		if len(positions) > 1 {
			for i := 0; i < len(positions)-1; i++ {
				if positions[i+1]-positions[i] > 1 {
					if positions[i+1]-positions[i] == 2 {
						if !(line[positions[i]] == line[positions[i+1]]) {
							appearstwice = true
						}
					}
					appearstwice = true
				}
			}
		}
	}

	for _, positions := range letters {
		if len(positions) > 1 {
			for i := 0; i < len(positions)-1; i++ {
				if positions[i+1]-positions[i] == 2 {
					oneletterinbetween = true
				}
			}
		}
	}

	if !appearstwice || !oneletterinbetween {
		return false
	}

	return true
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
