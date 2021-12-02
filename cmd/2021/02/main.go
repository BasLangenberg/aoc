package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}

	scanner := bufio.NewScanner(input)

	scanner.Split(bufio.ScanLines)

	hoz, depth := 0, 0
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		val, _ := strconv.Atoi(input[1])
		if input[0] == "forward" {
			hoz += val
		}
		if input[0] == "up" {
			depth -= val
		}
		if input[0] == "down" {
			depth += val
		}
	}

	fmt.Printf("Part 1: %d\n", hoz*depth)
}

func part2() {
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}

	scanner := bufio.NewScanner(input)

	scanner.Split(bufio.ScanLines)

	hoz, depth, aim := 0, 0, 0
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		val, _ := strconv.Atoi(input[1])
		if input[0] == "forward" {
			hoz += val
			depth += val * aim
		}
		if input[0] == "up" {
			aim -= val
		}
		if input[0] == "down" {
			aim += val
		}
	}

	fmt.Printf("Part 1: %d\n", hoz*depth)
}
