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

type bingocard struct {
	card   [][]bingonumber
	winner bool
}

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
		card := bingocard{}
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
			card.card = append(card.card, row)
		}
		bingocards = append(bingocards, card)
	}

	for _, card := range bingocards {
		if !validateCard(card) {
			fmt.Printf("ERROR Validating card! :%v", card)
			os.Exit(1)
		}
	}

	for iter, val := range bingonums {

		// The most inefficient code you've ever seen maybe
		var losers int

		for _, card := range bingocards {
			if !card.winner {
				losers++
			}
		}

		// fmt.Println(losers)

		if losers == 1 {
			for i, card := range bingocards {
				if !card.winner {
					fmt.Printf("Last winner: %d\n", i)
				}
			}
		}

		// Mark cards
		for i, card := range bingocards {
			markcard(card, val)
			if winner(&bingocards[i]) {
				fmt.Printf("Winner!, iter: %d, pos: %d, val: %d\n", iter, i, calcwin(card, val))
			}
		}

	}
}

func markcard(c bingocard, num int) {
	for _, row := range c.card {
		for x, item := range row {
			if item.num == num {
				(&row[x]).crossed = true
			}
		}
	}
}

func calcwin(c bingocard, num int) int {
	var sum int
	for _, row := range c.card {
		for _, val := range row {
			if !val.crossed {
				sum += val.num
			}
		}
	}

	return sum * num
}

func winner(c *bingocard) bool {
	// Can't win twice
	if c.winner {
		return false
	}

	// Check rows
	for _, row := range c.card {
		isWinner := true
		for _, item := range row {
			if !item.crossed {
				isWinner = false
				break
			}
		}

		if isWinner {
			c.winner = true
			return true
		}
	}
	// Check colums
	for i := 0; i < 5; i++ {
		isWinner := true
		for _, row := range c.card {
			if !row[i].crossed {
				isWinner = false
				break
			}
		}

		if isWinner {
			c.winner = true
			return true
		}
	}

	return false
}

func validateCard(c bingocard) bool {
	for _, row := range c.card {
		if len(row) != 5 {
			return false
		}
	}

	return true
}
