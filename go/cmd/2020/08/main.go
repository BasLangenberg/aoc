package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/BasLangenberg/aoc/internal/gameconsole"
)

func main() {
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var program []string

	for scanner.Scan() {
		program = append(program, scanner.Text())
	}

	compute := gameconsole.New(program)

	for {
		done := compute.Pulse()
		if done {
			fmt.Printf("ACU: %v, Pointer: %v\n", compute.Acu, compute.Pointer)
			break
		}
		if containsPreviousInstruction(compute.Pointer, compute.Previouspointers) {
			fmt.Printf("ACU: %v, Pointer: %v\n", compute.Acu, compute.Pointer)
			break
		}
	}

	for inst := range program {
		copyProgram := make([]string, len(program))
		copy(copyProgram, program)

		line := strings.Split(copyProgram[inst], " ")

		switch line[0] {
		case "jmp":
			copyProgram[inst] = strings.Replace(copyProgram[inst], "jmp", "nop", -1)
		case "nop":
			copyProgram[inst] = strings.Replace(copyProgram[inst], "nop", "jmp", -1)
		default:
			continue
		}

		comp := gameconsole.New(copyProgram)
		for {
			done := comp.Pulse()
			if done {
				fmt.Printf("ACU: %v, Pointer: %v\n", comp.Acu, comp.Pointer)
				break
			}
			if containsPreviousInstruction(comp.Pointer, comp.Previouspointers) {
				//fmt.Printf("looping after instruction %v\n", inst)
				break
			}
		}
	}

}

func containsPreviousInstruction(pointer int, previnstruct []int) bool {
	for _, ptr := range previnstruct {
		if pointer == ptr {
			return true
		}
	}

	return false
}
