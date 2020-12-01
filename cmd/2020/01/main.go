package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var expenses []int
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Unable to parse to int: %v", err)
			os.Exit(1)
		}

		expenses = append(expenses, num)
	}

	for i := 0; i < len(expenses); i++ {
		for y := 1; y < len(expenses); y++ {
			if expenses[i]+expenses[y] == 2020 {
				fmt.Printf("Part 1 solution: %v\n", expenses[i]*expenses[y])
			}
			for z := 2; z < len(expenses); z++ {
				if expenses[i]+expenses[y]+expenses[z] == 2020 {
					fmt.Printf("Part 2 solution: %v\n", expenses[i]*expenses[y]*expenses[z])
				}
			}
		}
	}
}
