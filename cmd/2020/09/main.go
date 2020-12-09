package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var inputnumbers []int

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Printf("Unable to parse as int: %v", err)
		}

		inputnumbers = append(inputnumbers, num)
	}

	preamble, values := getPreAmble(5, inputnumbers)

	found := make(map[int]bool)

	for pos, pre := range preamble {
		for _, num := range preamble[:pos] {
			found[pre+num] = true
		}
		for _, num := range preamble[pos+1:] {
			found[pre+num] = true
		}
	}
	fmt.Printf("%+v", found)

	for _, num := range values {
		if _, ok := found[num]; !ok {
			fmt.Println(num)
			break
		}
	}

}

func getPreAmble(preamb int, input []int) (preamble, values []int) {
	preamble = input[:preamb]
	values = input[preamb:]

	return preamble, values
}
