package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type x struct {
	zero, one int
}

func main() {
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var in []string

	for scanner.Scan() {
		in = append(in, scanner.Text())
	}

	var gamma string
	var epsilon string

	for i := 0; i < 12; i++ {
		var cur x
		for _, val := range in {
			if string(val[i]) == "0" {
				cur.zero++
			}
			if string(val[i]) == "1" {
				cur.one++
			}
		}

		if cur.zero > cur.one {
			gamma += "0"
			epsilon += "1"
		}

		if cur.one > cur.zero {
			gamma += "1"
			epsilon += "0"
		}
	}

	gam, _ := strconv.ParseInt(gamma, 2, 64)
	eps, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Printf("Part 1: %d\n", gam*eps)

	// Part 2
	// oxygen generator rating
	oxyinput := make([]string, len(in))
	copy(oxyinput, in)

	for i := 0; i < 12; i++ {
		zero, one := 0, 0
		if len(oxyinput) > 1 {
			for _, val := range oxyinput {
				if val[i] == '0' {
					zero++
				}
				if val[i] == '1' {
					one++
				}
			}

			if zero > one {
				oxyinput = returnBytePos(oxyinput, i, '0')
			}

			if one >= zero {
				oxyinput = returnBytePos(oxyinput, i, '1')
			}
		}
	}

	fmt.Println(oxyinput)

	// oxygen generator rating
	co2input := make([]string, len(in))
	copy(co2input, in)

	for i := 0; i < 12; i++ {
		zero, one := 0, 0
		if len(co2input) > 1 {
			for _, val := range co2input {
				if val[i] == '0' {
					zero++
				}
				if val[i] == '1' {
					one++
				}
			}

			if one < zero {
				co2input = returnBytePos(co2input, i, '1')
			}

			if zero <= one {
				co2input = returnBytePos(co2input, i, '0')
			}
		}
	}

	fmt.Println(co2input)

	ox, _ := strconv.ParseInt(oxyinput[0], 2, 64)
	co2, _ := strconv.ParseInt(co2input[0], 2, 64)

	fmt.Printf("Part 2: %d\n", ox*co2)
}

func returnBytePos(input []string, pos int, x byte) []string {
	var ret []string

	for _, val := range input {
		if val[pos] == x {
			ret = append(ret, val)
		}
	}

	return ret
}
