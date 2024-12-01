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

	total := 0

	for scanner.Scan() {
		count := 0

		ticket := scanner.Text()
		nums := strings.Split(ticket, ":")
		lotnums := strings.Split(nums[1], "|")
		lotnum := strings.Split(lotnums[0], " ")
		checknums := strings.Split(lotnums[1], " ")

		for _, nm := range lotnum {
			_, err := strconv.Atoi(nm)
			if err != nil {
				continue
			}
			for _, cnm := range checknums {
				if nm == cnm {
					fmt.Println(nm, cnm)
					count = updateCount(count)
				}
			}
		}
		total += count
	}
	fmt.Printf("Part 1: %v\n", total)
}

func updateCount(count int) int {
	if count == 0 {
		return 1
	}
	return count * 2
}
