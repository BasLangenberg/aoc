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
	var paper int
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		out, err := wrappingrequired(scanner.Text())
		if err != nil {
			fmt.Printf("Unable to read line: %v", err)
			os.Exit(1)
		}
		paper += out
	}

	fmt.Printf("Required paper: %d\n", paper)

}

func wrappingrequired(dimensions string) (int, error) {
	var dims []int
	var papersize int

	for _, size := range strings.Split(dimensions, "x") {
		dim, err := strconv.Atoi(size)
		if err != nil {
			return 0, fmt.Errorf("Unable to parse string to int: %v", size)
		}
		dims = append(dims, dim)
	}

	papersize += getsmallestdim(dims)
	papersize += 2 * (dims[0] * dims[1])
	papersize += 2 * (dims[0] * dims[2])
	papersize += 2 * (dims[1] * dims[2])

	return papersize, nil
}

func getsmallestdim(nums []int) int {
	smallest := math.MaxInt32

	var dims []int

	dims = append(dims, nums[0]*nums[1])
	dims = append(dims, nums[0]*nums[2])
	dims = append(dims, nums[1]*nums[2])
	for _, num := range dims {
		if num < smallest {
			smallest = num
		}
	}

	return smallest
}
