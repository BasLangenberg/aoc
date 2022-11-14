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
	var lint int

	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		out, err := wrappingrequired(line)
		if err != nil {
			fmt.Printf("Unable to read line: %v", err)
			os.Exit(1)
		}
		lout, err := lintrequired(line)
		if err != nil {
			fmt.Printf("Unable to read line: %v", err)
			os.Exit(1)
		}
		paper += out
		lint += lout
	}

	fmt.Printf("Required paper: %d\n", paper)
	fmt.Printf("Required lint: %d\n", lint)

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

func lintrequired(dimensions string) (int, error) {
	var dims []int
	var smalldims []int
	var lintsize int

	for _, size := range strings.Split(dimensions, "x") {
		dim, err := strconv.Atoi(size)
		if err != nil {
			return 0, fmt.Errorf("Unable to parse string to int: %v", size)
		}
		dims = append(dims, dim)
	}

	lintsize += dims[0] * dims[1] * dims[2]

	for i := 0; i < 2; i++ {
		small := getsmallestnum(dims)
		for pos, val := range dims {
			if val == small {
				smalldims = append(smalldims, small)
				dims[len(dims)-1], dims[pos] = dims[pos], dims[len(dims)-1]
				dims = dims[:len(dims)-1]
				break
			}
		}
	}

	if len(smalldims) == 1 {
		smalldims[1] = smalldims[0]
	}

	lintsize += smalldims[0] + smalldims[0] + smalldims[1] + smalldims[1]
	return lintsize, nil
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

// TODO Refactor to own package
func getsmallestnum(input []int) int {
	resp := math.MaxInt32
	for _, i := range input {
		if i < resp {
			resp = i
		}
	}
	return resp
}
