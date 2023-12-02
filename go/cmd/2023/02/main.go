package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("part 1: %v\n", part1())
	fmt.Printf("part 2: %v\n", part2())
}

func part1() int {
	check := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	gamecounter := 0
	for scanner.Scan() {
		valid := true
		splt := strings.Split(scanner.Text(), ":")
		game, err := strconv.Atoi(strings.Split(splt[0], " ")[1])
		if err != nil {
			panic("Unable to get game number")
		}

		cubecounter := make(map[string]int)

		line := splt[1]
		line = strings.Replace(line, ";", ",", -1)

		for _, v := range strings.Split(line, ",") {
			itm := strings.Split(strings.Trim(v, " "), " ")
			fmt.Println(itm)
			num, err := strconv.Atoi(itm[0])
			if err != nil {
				panic("Unable to parse cube count number")
			}

			col := itm[1]
			cubec := cubecounter[col]
			if cubec < num {
				cubecounter[col] = num
			}
		}
		for k, v := range cubecounter {
			if check[k] < v {
				valid = false
			}
		}

		if valid {
			gamecounter += game
		}
		fmt.Println(cubecounter)
	}

	return gamecounter

}

func part2() int {
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	gamecounter := 0
	for scanner.Scan() {
		splt := strings.Split(scanner.Text(), ":")
	//	game, err := strconv.Atoi(strings.Split(splt[0], " ")[1])
		if err != nil {
			panic("Unable to get game number")
		}

		cubecounter := make(map[string]int)

		line := splt[1]
		line = strings.Replace(line, ";", ",", -1)

		for _, v := range strings.Split(line, ",") {
			itm := strings.Split(strings.Trim(v, " "), " ")
			fmt.Println(itm)
			num, err := strconv.Atoi(itm[0])
			if err != nil {
				panic("Unable to parse cube count number")
			}

			col := itm[1]
			cubec := cubecounter[col]
			if cubec < num {
				cubecounter[col] = num
			}
		}
		gamecounter += cubecounter["red"] * cubecounter["blue"] * cubecounter["green"]
	}

	return gamecounter

}

