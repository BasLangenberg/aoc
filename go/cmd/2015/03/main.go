package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type pos struct {
	horizontal int
	vertical   int
}

func main() {
	input, err := ioutil.ReadFile("input")

	if err != nil {
		fmt.Println("Unable to find file called input in local directory")
		os.Exit(1)
	}

	fmt.Printf("Houses visited: %d\n", housesvisited(string(input)))
	fmt.Printf("Houses visited with robot: %d\n", housesvisitedwrobotsanta(string(input)))

}
func housesvisited(directions string) int {
	set := make(map[string]bool)
	santapos := pos{
		horizontal: 0,
		vertical:   0,
	}

	set["0,0"] = true

	for _, dir := range directions {
		switch dir {
		case '>':
			santapos.vertical = santapos.vertical + 1
		case '<':
			santapos.vertical = santapos.vertical - 1
		case '^':
			santapos.horizontal = santapos.horizontal + 1
		case 'v':
			santapos.horizontal = santapos.horizontal - 1
		}

		set[strconv.Itoa(santapos.horizontal)+","+strconv.Itoa(santapos.vertical)] = true

	}
	return len(set)
}

func housesvisitedwrobotsanta(directions string) int {
	santaset := make(map[string]bool)
	roboset := make(map[string]bool)

	santapos := pos{
		horizontal: 0,
		vertical:   0,
	}

	robopos := pos{
		horizontal: 0,
		vertical:   0,
	}

	santaset[strconv.Itoa(santapos.horizontal)+","+strconv.Itoa(santapos.vertical)] = true
	roboset[strconv.Itoa(robopos.horizontal)+","+strconv.Itoa(robopos.vertical)] = true

	for instr, dir := range directions {
		switch instr % 2 {
		case 0:
			switch dir {
			case '>':
				santapos.vertical = santapos.vertical + 1
			case '<':
				santapos.vertical = santapos.vertical - 1
			case '^':
				santapos.horizontal = santapos.horizontal + 1
			case 'v':
				santapos.horizontal = santapos.horizontal - 1
			}
			santaset[strconv.Itoa(santapos.horizontal)+","+strconv.Itoa(santapos.vertical)] = true

		default:
			switch dir {
			case '>':
				robopos.vertical = robopos.vertical + 1
			case '<':
				robopos.vertical = robopos.vertical - 1
			case '^':
				robopos.horizontal = robopos.horizontal + 1
			case 'v':
				robopos.horizontal = robopos.horizontal - 1
			}

			roboset[strconv.Itoa(robopos.horizontal)+","+strconv.Itoa(robopos.vertical)] = true
		}
	}
	houses := len(santaset)
	for key := range roboset {
		if _, ok := santaset[key]; !ok {
			houses++
		}
	}

	return houses
}
