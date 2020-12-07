package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bag struct {
	name   string
	amount int
	bags   []bag
}

func main() {
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}

	bags := []bag{}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		nb := bag{
			name:   strings.Split(line, " ")[0] + " " + strings.Split(line, " ")[1] + " bag",
			amount: 1,
		}

		containingBags := strings.Join(strings.Split(line, " ")[4:], " ")
		containingBags = strings.Replace(containingBags, ", ", ",", -1)

		for _, bg := range strings.Split(containingBags, ",") {
			if bg[0] == 'n' {
				continue
			}

			amount, err := strconv.Atoi(string(bg[0]))
			if err != nil {
				fmt.Printf("Unable to parse as int: %v\n", err)
			}
			b := bag{
				name:   fmt.Sprintf("%v bag", strings.Join(strings.Split(bg, " ")[1:3], " ")),
				amount: amount,
			}

			nb.bags = append(nb.bags, b)
		}
		bags = append(bags, nb)
	}
	foundBags := findContains(bags, "shiny gold bag")
	uniqFoundBags := make(map[string]bool)

	for _, bag := range foundBags {
		uniqFoundBags[bag] = true
	}

	fmt.Println(len(uniqFoundBags))
	fmt.Println(bagsInbag("shiny gold bag", bags))

}

func findContains(bags []bag, bag string) []string {
	var found []string
	for _, bg := range bags {
		for _, b := range bg.bags {
			if b.name == bag {
				found = append(found, bg.name)
				found = append(found, findContains(bags, bg.name)...)
			}
		}
	}

	return found
}

func bagsInbag(bag string, bags []bag) int {
	var count int

	for _, bg := range bags {
		if bg.name == bag {
			if len(bg.bags) == 0 {
				return 0
			}
			for _, b := range bg.bags {
				inbags := bagsInbag(b.name, bags)
				count += (b.amount * inbags) + b.amount
			}
		}
	}
	return count
}
