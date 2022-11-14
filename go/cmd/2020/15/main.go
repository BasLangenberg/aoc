package main

import (
	"bufio"
	"fmt"
	"os"
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
	var memo [][]int

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		var iLine []int

		for _, num := range line {
			nm, err := strconv.Atoi(num)

			if err != nil {
				fmt.Printf("unable to parse int: %v", err)
				os.Exit(1)
			}

			iLine = append(iLine, nm)
		}

		memo = append(memo, iLine)

	}

	// Part A
	for p, seq := range memo {
		memos := make(map[int][]int)
		for pos, num := range seq {
			memos[num] = append(memos[num], pos+1)
		}

		for i := len(memos); i < 2020; i++ {
			oldpos := seq[len(seq)-1]
			if len(memos[oldpos]) == 1 {
				memos[0] = append(memos[0], i+1)
				seq = append(seq, 0)
				continue
			}

			occur := memos[oldpos]
			val := occur[len(occur)-1] - occur[len(occur)-2]
			memos[val] = append(memos[val], i+1)
			seq = append(seq, val)
		}

		fmt.Printf("Value 2020 for input line %+v is %+v\n", p, seq[len(seq)-1])
		// fmt.Printf("Sequence: %+v", seq)
	}

	// Part B
	for p, seq := range memo {
		memos := make(map[int][]int)
		for pos, num := range seq {
			memos[num] = append(memos[num], pos+1)
		}

		for i := len(memos); i < 30000000; i++ {
			oldpos := seq[len(seq)-1]
			if len(memos[oldpos]) == 1 {
				memos[0] = append(memos[0], i+1)
				seq = append(seq, 0)
				continue
			}

			occur := memos[oldpos]
			val := occur[len(occur)-1] - occur[len(occur)-2]
			memos[val] = append(memos[val], i+1)
			seq = append(seq, val)
		}

		fmt.Printf("Value 2020 for input line %+v is %+v\n", p, seq[len(seq)-1])
		// fmt.Printf("Sequence: %+v", seq)
	}
}
