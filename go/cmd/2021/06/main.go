package main

import (
	"bufio"
	"fmt"
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

	var fish []int

	for scanner.Scan() {
		for _, val := range strings.Split(scanner.Text(), ",") {
			f, _ := strconv.Atoi(val)
			fish = append(fish, f)
		}
	}

	eightydays := simulateRepro(fish, 80)
	fmt.Printf("Number of fish after 80 days: %d\n", eightydays)

	twohondredsixtyfive := simulateRepro(fish, 256)
	fmt.Printf("Number of fish after 256 days: %d\n", twohondredsixtyfive)

}

func simulateRepro(fish []int, days int) int {
	school := make(map[int]int)

	for i := 0; i < 9; i++ {
		school[i] = 0
	}

	for _, val := range fish {
		school[val] += 1
	}

	for i := 0; i < days; i++ {
		tmp := make(map[int]int)
		tmp[0] = school[1]
		tmp[1] = school[2]
		tmp[2] = school[3]
		tmp[3] = school[4]
		tmp[4] = school[5]
		tmp[5] = school[6]
		tmp[6] = school[7] + school[0]
		tmp[7] = school[8]
		tmp[8] = school[0]

		school = tmp
	}

	var ret int
	for _, val := range school {
		ret += val
	}

	return ret
}
