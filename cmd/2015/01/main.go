package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Get input
	input, err := ioutil.ReadFile("input")

	if err != nil {
		fmt.Println("Unable to find file called input in local directory")
		os.Exit(1)
	}

	fmt.Printf("Final floor: %d", getfloor(input))
}
func getfloor(instructions []byte) int {
	var floor int
	for _, char := range instructions {
		if string(char) == ")" {
			floor--
		}
		if string(char) == "(" {
			floor++
		}
	}
	return floor
}
