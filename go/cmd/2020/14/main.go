package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	var program []string

	for scanner.Scan() {
		program = append(program, strings.ReplaceAll(scanner.Text(), " ", ""))
	}

	var mask string
	var addr int
	var val int
	mem := make(map[int]int)
	r := regexp.MustCompile("\\d+")

	for _, inst := range program {
		i := strings.Split(inst, "=")
		if i[0] == "mask" {
			mask = i[1]
			continue
		}

		if i[0][:3] == "mem" {
			addr, err = strconv.Atoi(string(r.Find([]byte(i[0]))))
			if err != nil {
				fmt.Printf("Unable to parse int: %v", err)
				os.Exit(1)
			}

			val, err = strconv.Atoi(i[1])
			if err != nil {
				fmt.Printf("Unable to parse int: %v", err)
				os.Exit(1)
			}

			var prepend string
			binval := strconv.FormatInt(int64(val), 2)

			for i := 0; i < 36-len(binval); i++ {
				prepend += "0"
			}

			binval = prepend + binval
			fmt.Printf("M: %v\n", mask)
			fmt.Printf("I: %v\n", binval)
			rbinval := []rune(binval)

			for i := len(binval) - 1; i >= 0; i-- {
				if mask[i] == '1' {
					rbinval[i] = '1'
				}
				if mask[i] == '0' {
					rbinval[i] = '0'
				}
			}

			binval = string(rbinval)
			fmt.Printf("O: %v\n", binval)

			realvalue, err := strconv.ParseInt(binval, 2, 64)
			if err != nil {
				fmt.Printf("Unable to parse int: %v", err)
				os.Exit(1)
			}

			mem[addr] = int(realvalue)
		}
	}

	var total int
	for _, num := range mem {
		total += num
	}

	fmt.Printf("Too low: %v\n", 237592804102)
	fmt.Printf("Total: %+v\n", total)
	fmt.Printf("%+v", mem)
}
