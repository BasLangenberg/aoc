package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bingonumber struct {
	num     int
	crossed bool
}

type bingocard [][]bingonumber

func main() {
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	// Process first line
	scanner.Scan()

	bingo := strings.Split(scanner.Text(), ",")

	var bingonums []int
	for _, item := range bingo {
		num, _ := strconv.Atoi(item)
		bingonums = append(bingonums, num)
	}

	// Create bingocards
	var bingocards []bingocard

	for scanner.Scan() {
		var card bingocard
		for i := 0; i < 5; i++ {
			var row []bingonumber
			scanner.Scan()
			l := scanner.Text()
			if l[0] == ' ' {
				l = l[1:]
			}
			l = strings.Replace(l, "  ", " ", -1)
			for _, val := range strings.Split(l, " ") {
				num, _ := strconv.Atoi(val)
				item := bingonumber{num, false}
				row = append(row, item)
			}
			card = append(card, row)
		}
		bingocards = append(bingocards, card)
	}

	for _, card := range bingocards {
		if !validateCard(card) {
			fmt.Printf("ERROR Validating card! :%v", card)
			os.Exit(1)
		}
	}

	for _, val := range bingonums {
		// Mark cards
		for _, card := range bingocards {
			markcard(card, val)
		}
		// Check for winners
		for _, card := range bingocards {
			if winner(card) {
				fmt.Printf("Part 1: %d", calcwin(card, val))
				os.Exit(0)
			}

		}
	}
}

func markcard(c bingocard, num int) {
	for _, row := range c {
		for x, item := range row {
			if item.num == num {
				(&row[x]).crossed = true
			}
		}
	}
}

func calcwin(c bingocard, num int) int {
	var sum int
	for _, row := range c {
		for _, val := range row {
			if !val.crossed {
				sum += val.num
			}
		}
	}

	return sum * num
}

func winner(c bingocard) bool {
	// Check rows
	for _, row := range c {
		isWinner := true
		for _, item := range row {
			if !item.crossed {
				isWinner = false
				break
			}
		}

		if isWinner {

			return true
		}
	}
	// Check colums
	for i := 0; i < 5; i++ {
		isWinner := true
		for _, row := range c {
			if !row[i].crossed {
				isWinner = false
				break
			}
		}

		if isWinner {
			return true
		}
	}

	return false
}

func validateCard(c bingocard) bool {
	for _, row := range c {
		if len(row) != 5 {
			return false
		}
	}

	return true
}
