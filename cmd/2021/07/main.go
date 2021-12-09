package main

import (
	"bufio"
	"fmt"
	"math"
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
	scanner.Scan()
	var inp []int
	nums := strings.Split(scanner.Text(), ",")
	for _, val := range nums {
		vl, _ := strconv.Atoi(val)
		inp = append(inp, vl)
	}

	fmt.Printf("Part 1: %d\n", leastfuel(inp))
	fmt.Printf("Part 2: %d\n", leastfuel2(inp))

}

func leastfuel(input []int) int {
	fuel := math.MaxInt64
	for _, val := range input {
		var f int
		for _, vl := range input {
			f += Abs(vl - val)
		}
		if fuel > f {
			fuel = f
		}
	}
	return fuel
}

func leastfuel2(input []int) int {
	fuel := math.MaxInt64
	maxNum := getmax(input)
	minNum := getmin(input)

	for i := minNum; i <= maxNum; i++ {
		var f int
		for _, vl := range input {
			x := Abs(vl - i)
			for i := 1; i <= x; i++ {
				f += i
			}
		}
		if fuel > f {
			fuel = f
		}
	}
	return fuel
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getmax(i []int) int {
	min := math.MinInt64
	for _, n := range i {
		if n > min {
			min = n
		}
	}
	return min
}

func getmin(i []int) int {
	max := math.MaxInt64
	for _, n := range i {
		if n < max {
			max = n
		}
	}
	return max
}
