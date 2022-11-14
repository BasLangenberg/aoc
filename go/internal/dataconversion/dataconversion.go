// Package dataconversion is written for use with AOC problems only
// Therefore, I do not think it is an antipattern to terminate inside this package
// Huge exception to normal coding standards!
package dataconversion

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// StringToIntArray takes a comma seperated string of numbers and converts it to an array of ints
func StringToIntArray(input string) []int {
	var arr []int
	for _, num := range strings.Split(input, ",") {
		n, err := strconv.Atoi(num)
		if err != nil {
			fmt.Printf("unable to parse int: %v", err)
			os.Exit(1)
		}
		arr = append(arr, n)

	}

	return arr
}

// IntInSlice checks if an integer exists in a slice of int
func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
