package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func main() {
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var plan [][]rune

	for scanner.Scan() {
		var line []rune
		for _, pos := range scanner.Text() {
			line = append(line, pos)
		}
		plan = append(plan, line)
	}

	//	fmt.Printf("%+v", plan)
	cache := make(map[int][][]rune)
	count := 0

	cache[count] = plan

	for {
		fmt.Printf("Pass: %v\n", count)

		newplan := make([][]rune, len(cache[count]))

		for i := range cache[count] {
			newplan[i] = make([]rune, len(cache[count][i]))
			copy(newplan[i], cache[count][i])
		}

		newplan = parseMap(newplan)

		printMap(cache[count])
		printMap(newplan)

		if reflect.DeepEqual(newplan, cache[count]) {
			break
		}

		count++
		cache[count] = newplan
	}

	countOccupied := 0

	for _, line := range cache[count] {
		for _, char := range line {
			if char == '#' {
				countOccupied++
			}
		}
	}

	fmt.Printf("Seats occupied: %v\n", countOccupied)
}

func printMap(plan [][]rune) {
	fmt.Println()
	fmt.Println("===================================================")
	fmt.Println()
	for _, i := range plan {
		fmt.Println(string(i))
	}
}

func parseMap(plan [][]rune) [][]rune {
	newplan := make([][]rune, len(plan))

	for i := range plan {
		newplan[i] = make([]rune, len(plan[i]))
		copy(newplan[i], plan[i])
	}

	for vpos, line := range plan {
		for hpos, char := range line {
			if char == '.' {
				continue
			}
			occupied, empty := parseNeighbours(hpos, vpos, plan)
			fmt.Printf("Cord: %v,%v - occupied: %v, empty: %v\n", hpos, vpos, occupied, empty)
			if plan[vpos][hpos] == 'L' {
				if occupied == 0 {
					newplan[vpos][hpos] = '#'
				}
			}
			if plan[vpos][hpos] == '#' {
				if occupied > 3 {
					newplan[vpos][hpos] = 'L'
				}
			}
		}
	}

	return newplan
}

func parseNeighbours(hpos, vpos int, plan [][]rune) (int, int) {
	empty := 0
	occupied := 0

	// Links
	if hpos > 0 {
		if plan[vpos][hpos-1] == 'L' {
			empty++
		}
		if plan[vpos][hpos-1] == '#' {
			occupied++
		}
	}

	// Linksboven
	if vpos > 0 && hpos > 0 {
		if plan[vpos-1][hpos-1] == 'L' {
			empty++
		}
		if plan[vpos-1][hpos-1] == '#' {
			occupied++
		}
	}

	// Boven
	if vpos > 0 {
		if plan[vpos-1][hpos] == 'L' {
			empty++
		}
		if plan[vpos-1][hpos] == '#' {
			occupied++
		}
	}

	// Rechtsboven
	if hpos < len(plan[vpos])-1 && vpos > 0 {
		if plan[vpos-1][hpos+1] == 'L' {
			empty++
		}
		if plan[vpos-1][hpos+1] == '#' {
			occupied++
		}
	}

	// Rechts
	if hpos < len(plan[vpos])-1 {
		if plan[vpos][hpos+1] == 'L' {
			empty++
		}
		if plan[vpos][hpos+1] == '#' {
			occupied++
		}
	}

	// Rechtsonder
	if vpos < len(plan)-1 && hpos < len(plan[vpos])-1 {
		if plan[vpos+1][hpos+1] == 'L' {
			empty++
		}
		if plan[vpos+1][hpos+1] == '#' {
			occupied++
		}
	}

	// Onder
	if vpos < len(plan)-1 {
		if plan[vpos+1][hpos] == 'L' {
			empty++
		}
		if plan[vpos+1][hpos] == '#' {
			occupied++
		}
	}

	// linksonder
	if vpos > 0 && hpos > 0 {
		if plan[vpos-1][hpos-1] == 'L' {
			empty++
		}
		if plan[vpos-1][hpos-1] == '#' {
			occupied++
		}
	}

	return occupied, empty
}
