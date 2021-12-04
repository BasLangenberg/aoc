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
	// Check
	// fmt.Printf("Input bingo: %s\n\n", bingo)
	// for _, val := range bingocards {
	// for _, row := range val {
	// for _, num := range row {
	// fmt.Printf("%d ", num.num)
	// }
	// fmt.Printf("\n")
	// }
	// fmt.Printf("\n")
	// }

	var windex []int
	for iter, val := range bingonums {
		// Mark cards
		for _, card := range bingocards {
			markcard(card, val)
		}
		// Check for winners
		for k, card := range bingocards {
			// if winner(card) && len(bingocards) == 1 {
			// 	fmt.Printf("Part 2: %d - iteration: %d\n", calcwin(card, val), iter)
			// 	os.Exit(0)
			// }

			// if winner(card) {
			// 	if len(bingocards) == k {
			// 		bingocards = bingocards[:k]
			// 		continue
			// 	}
			// 	bingocards = append(bingocards[:k], bingocards[k+1:]...)
			// }
			if winner(card) {
				windex = append(windex, k)
			}

			if len(windex) == 99 {
				for i := 0; i < 100; i++ {
					found := true
					for _, val := range windex {
						if val == i {
							found = false
						}
					}

					if found {
						calcwin(bingocards[i], iter)
						os.Exit(0)
					}
				}
			}

		}
		fmt.Printf("Iter: %d - Len: %d\n", iter, len(windex))
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
