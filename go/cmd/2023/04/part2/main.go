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

	tickets := make(map[int][][]int)
	check := make(map[int][]int)
	tcknum := 0

	for scanner.Scan() {
		tcknum += 1

		ticket := scanner.Text()
		nums := strings.Split(ticket, ":")
		lotnums := strings.Split(nums[1], "|")
		lotnum := strings.Split(lotnums[0], " ")
		checknums := strings.Split(lotnums[1], " ")

		lotnms := []int{}
		for _, nm := range lotnum {
			n, err := strconv.Atoi(nm)
			if err != nil {
				continue
			}
			lotnms = append(lotnms, n)
		}

		chknms := []int{}
		for _, nm := range checknums {
			n, err := strconv.Atoi(nm)
			if err != nil {
				continue
			}
			chknms = append(chknms, n)
		}

		tickets[tcknum] = append(tickets[tcknum], lotnms)
		check[tcknum] = chknms
	}

	for i := 1; i < len(tickets); i++ {
		counter := 0

		for _, tick := range tickets[i] {
			for _, nmbr := range tick {
				for _, chk := range check[i] {
					if nmbr == chk {
						counter += 1
						if counter+i > len(tickets) {
							continue
						}
						cpy := tickets[i+counter][0]
						tickets[i+counter] = append(tickets[i+counter], cpy)
					}
				}
			}
		}
	}

	cnt := 0
	for k, v := range tickets {
		fmt.Printf("ticket %v: %v\n", k, len(v))

		cnt += len(v)
	}

	fmt.Printf("Part 2: %v\n", cnt)
}

func updateCount(count int) int {
	return count + 1
}
