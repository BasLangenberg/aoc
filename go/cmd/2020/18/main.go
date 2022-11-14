package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		// CODE GOES HERE
	}

}

func calc(input string) string {
	regexp.MustCompile(`\(.*\)`)

}
