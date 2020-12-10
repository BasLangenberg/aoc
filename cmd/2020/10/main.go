package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var inputnumbers []int
	inputnumbers = append(inputnumbers, 0)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Printf("Unable to parse as int: %v", err)
		}

		inputnumbers = append(inputnumbers, num)
	}

	sort.Ints(inputnumbers)
	inputnumbers = append(inputnumbers, inputnumbers[len(inputnumbers)-1]+3)

	fmt.Printf("Adapters:\n%v\n", inputnumbers)

	diffOne := 0
	diffThree := 0

	for i := 1; i < len(inputnumbers); i++ {
		if inputnumbers[i]-inputnumbers[i-1] == 1 {
			diffOne++
		}
		if inputnumbers[i]-inputnumbers[i-1] == 3 {
			diffThree++
		}
	}

	fmt.Printf("Part A: %v x %v = %v\n", diffOne, diffThree, diffOne*diffThree)

	// Part B
	inputs := cut(inputnumbers)
	result := 1
	for _, inp := range inputs {
		result *= calc(1, inp)
		fmt.Printf("Piece: %v, Paths: %v\n", inp, calc(1, inp))
	}
	fmt.Printf("Too low (Part B): %v\n", 4294967296)
	fmt.Printf("Too low (Part B): %v\n", 2739820091932672)
	fmt.Printf("Part B: %v\n", result)

}

func cut(input []int) [][]int {
	var inputs [][]int
	prev := 0
	for i := 1; i < len(input); i++ {

		if input[i]-input[i-1] == 3 {
			inputs = append(inputs, input[prev:i])
			prev = i
		}
	}

	return inputs
}

func calc(pos int, adapters []int) int {
	possible := 1
	//	fmt.Printf("Start pos: %v, len: %v\n", pos, len(adapters))
	//	fmt.Println(len(adapters))
	for i := pos; i < len(adapters); i++ {
		addCount := 0
		switch adapters[i] - adapters[i-1] {
		case 3:
			continue
		case 2:
			if len(adapters) > i+1 {
				if adapters[i+1]-adapters[i-1] == 1 {
					addCount++
					possible += calc(i+1, adapters)
				}
			}
		case 1:
			if len(adapters) > i+1 {
				if adapters[i+1]-adapters[i-1] == 2 || adapters[i+1]-adapters[i-1] == 3 {
					addCount++
					possible += calc(i+1, adapters)
				}
			}

			if len(adapters) > i+2 {
				if adapters[i+2]-adapters[i-1] == 3 {
					addCount++
					possible += calc(i+2, adapters)
				}
			}
		default:
			fmt.Printf("Value out of range: %v, position: %v, Value: %v", adapters[i]-adapters[i-1], i, possible)
		}
		i += addCount
	}

	return possible
}
