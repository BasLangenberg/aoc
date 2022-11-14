package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

	type policy struct {
		min, max int
		letter   string
	}

	type password struct {
		passwd string
		valid  bool
		pol    policy
	}

	memo := make(map[string]password)
	var correct int

	for scanner.Scan() {
		line := scanner.Text()

		pass := strings.Split(line, ": ")
		let := strings.Split(line, " ")
		minmax := strings.Split(let[0], "-")

		minoc, err := strconv.Atoi(string(minmax[0]))
		if err != nil {
			fmt.Printf("Unable to parse minimum as int: %v", err)
		}

		maxoc, err := strconv.Atoi(string(minmax[1]))
		if err != nil {
			fmt.Printf("Unable to parse maximum as int: %v", err)
		}

		pw := password{
			passwd: pass[1],
			pol: policy{
				min:    minoc,
				max:    maxoc,
				letter: string(let[1][0]),
			},
		}

		reg := regexp.MustCompile(pw.pol.letter)
		occurrances := reg.FindAllStringIndex(pw.passwd, -1)

		if len(occurrances) > pw.pol.max || len(occurrances) < pw.pol.min {
			pw.valid = false
			memo[pw.passwd] = pw
			continue
		}

		pw.valid = true
		memo[pw.passwd] = pw
		correct++

	}

	fmt.Printf("Correct passwords: %v\n", correct)

	// Part 2
	var part2count int

	for _, pwd := range memo {
		var pos1 bool
		var pos2 bool

		if pwd.pol.letter == string(pwd.passwd[pwd.pol.min-1]) {
			pos1 = true
		}

		if pwd.pol.letter == string(pwd.passwd[pwd.pol.max-1]) {
			pos2 = true
		}

		if pos1 {
			if !pos2 {
				part2count++
			}
		}

		if !pos1 {
			if pos2 {
				part2count++
			}
		}
	}

	fmt.Printf("Correct passwords part 2: %v\n", part2count)
}
