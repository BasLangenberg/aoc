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

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Printf("Unable to parse as int: %v", err)
		}

		inputnumbers = append(inputnumbers, num)
	}

	pvalue := 25

	for i := pvalue; i < (len(inputnumbers) - pvalue); i++ {
		found := make(map[int]bool)

		for y := i - pvalue; y < i; y++ {
			for _, num := range inputnumbers[i-pvalue : y] {
				fnd := num + inputnumbers[y]
				found[fnd] = true
			}
			for _, num := range inputnumbers[y+1 : i] {
				fnd := num + inputnumbers[y]
				found[fnd] = true
			}
		}

		if _, ok := found[inputnumbers[i]]; !ok {
			fmt.Println(inputnumbers[i])
		}
	}

	toMatch := 756008079

	for i := 0; i < len(inputnumbers); i++ {
		var tmp int
		var list []int
		for y := i + 1; y < len(inputnumbers); y++ {
			if tmp > toMatch {
				break
			}

			tmp += inputnumbers[y]
			list = append(list, inputnumbers[y])

			if tmp == toMatch {
				fmt.Printf("Range found: %+v\n", list)
				sort.Ints(list)
				fmt.Printf("Answer: %v (lowest + highest)\n", list[0]+list[len(list)-1])
			}
		}
	}

}
