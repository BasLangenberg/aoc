package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input")
	data := []map[rune]int{}

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}

	groups := strings.Split(string(input), "\n\n")

	for _, group := range groups {
		parsed := make(map[rune]int)
		group = strings.ReplaceAll(group, "\n", "")
		for _, let := range group {
			if _, ok := parsed[let]; ok {
				parsed[let] = parsed[let] + 1
			}
			parsed[let] = 1
		}

		data = append(data, parsed)
	}

	var total int
	for _, group := range data {
		total += len(group)
	}

	fmt.Printf("Number of unique questions answered yes: %v\n", total)

	var allAnsweredYes int
	for _, group := range groups {
		numPeeps := len(strings.Split(group, "\n"))
		parsed := make(map[rune]int)
		group = strings.ReplaceAll(group, "\n", "")

		for _, let := range group {
			if num, ok := parsed[let]; ok {
				parsed[let] = num + 1
				continue
			}
			parsed[let] = 1
		}

		for _, count := range parsed {
			if count == numPeeps {
				allAnsweredYes++
			}
		}

	}

	fmt.Printf("Number of group questions answered yes: %v\n", allAnsweredYes)
}
