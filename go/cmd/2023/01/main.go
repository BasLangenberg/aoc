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
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var decs [][]string
	for scanner.Scan() {
		var dec []string
		for _, c := range scanner.Text() {
			_, err := strconv.Atoi(string(c))
			if err == nil {
				dec = append(dec, string(c))
			}
		}
		decs = append(decs, dec)
	}

	var count int
	for _, line := range decs {
		number := line[0] + line[len(line)-1]
		num, _ := strconv.Atoi(number)
		count += num
	}

	return count
}

func part2() int {
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var decs [][]string
	for scanner.Scan() {
		var dec []string
		line := scanner.Text()
		line = replace(line)

		for _, c := range line {
			_, err := strconv.Atoi(string(c))
			if err == nil {
				dec = append(dec, string(c))
			}
		}
		decs = append(decs, dec)
	}

	var count int
	for _, line := range decs {
		number := line[0] + line[len(line)-1]
		num, _ := strconv.Atoi(number)
		count += num
	}

	return count
}

func replace(inp string) string {
	integers := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	var first, last string
	for i := range inp {
		tmp := inp[i :]
		for k, v := range integers {
			if strings.HasPrefix(tmp, k) {
				first = v
				break
			}
		}
		if first != "" {
			break
		}
		_, err := strconv.Atoi(string(tmp[0]))
		if err == nil {
			first = string(tmp[0])
			break
		}
	}

	for i := len(inp)-1; i > 0; i-- {
		tmp := inp[i:]
		for k, v := range integers {
			if strings.Contains(tmp, k) {
				last = v
				break
			}
		}
		if last != "" {
			break
		}
//		fmt.Println(tmp)
		_, error := strconv.Atoi(string(tmp[0]))
		if error == nil {
			last = string(tmp[0])
			break
		}
	}
	return first+last
}
