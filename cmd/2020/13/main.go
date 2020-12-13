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
	scanner.Split(bufio.ScanLines)
	var data string

	for scanner.Scan() {
		data = scanner.Text()
	}

	parsed := strings.Split(data, ",")
	timestamp, err := strconv.Atoi(parsed[0])

	if err != nil {
		fmt.Printf("Unable to parse number: %v", err)
	}

	busses := make(map[int][]int)

	for i := 1; i < len(parsed); i++ {
		if parsed[i] == "x" {
			continue
		}

		num, err := strconv.Atoi(parsed[i])
		if err != nil {
			fmt.Printf("Unable to parse number: %v", err)
		}
		busses[num] = []int{0}
	}

	for key := range busses {
		val := 0
		for {
			if val > timestamp {
				break
			}

			val = val + key
			busses[key] = append(busses[key], val)
		}
	}

	var bus int
	timeTillDepart := math.MaxInt64

	for key, departts := range busses {
		if departts[len(busses[key])-1]-timestamp < timeTillDepart {
			timeTillDepart = departts[len(busses[key])-1] - timestamp
			bus = key
		}

	}

	fmt.Printf("Bus: %v, Departs in %v minutes\n", bus, timeTillDepart)
	fmt.Printf("Answer: %v\n", bus*timeTillDepart)

	// Part b

	input, err = os.Open("inputb")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}

	scanner = bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	var pb []string

	for scanner.Scan() {
		pb = append(pb, scanner.Text())
	}

	for line, val := range pb {
		fmt.Printf("line: %v - timestamp: %v\n", line, findSeq(val))
	}

}

func findSeq(line string) int {
	var data []int

	for _, num := range strings.Split(line, ",") {
		if num == "x" {
			data = append(data, 0)
			continue
		}

		// Never really ignore the error of course! :-)
		nm, _ := strconv.Atoi(num)
		data = append(data, nm)
	}

	t := 0

	busses := make(map[int][]int)

	// initialize the list
	for _, num := range data {
		if num == 0 {
			continue
		}
		busses[num] = []int{num}
	}

	for {
		checkds := make(map[int]bool)
		for key := range busses {
			if busses[key][len(busses[key])-1] < t {
				busses[key] = append(busses[key], busses[key][len(busses[key])-1]+key)
			}
			checkds[key] = false
		}

		// Validation
		counter := 0
		for _, val := range data {
			if val == 0 {
				counter++
				continue
			}

			if busses[val][len(busses[val])-1] == t+counter {
				checkds[val] = true
			}
			counter++
		}

		allOk := true
		for _, val := range checkds {
			if !val {
				allOk = false
				break
			}
		}

		if allOk {
			return t
		}

		// Increase T
		t++
	}
}
