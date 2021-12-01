package main

import (
	"bufio"
	"fmt"
	"os"
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

	var numbers []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Errorf("Unable to parse string: %s", err)
			os.Exit(1)
		}

		numbers = append(numbers, num)
	}

	var count int

	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] {
			count++
		}
	}

	fmt.Printf("Part 1: %d\n", count)

	var slidingWindow []int
	for i := 2; i < len(numbers); i++ {
		slidingWindow = append(slidingWindow, (numbers[i-2] + numbers[i-1] + numbers[i]))
	}

	var countp2 int
	for i := 1; i < len(slidingWindow); i++ {
		if slidingWindow[i] > slidingWindow[i-1] {
			countp2++
		}
	}

	fmt.Printf("Part 2: %d\n", countp2)

}
