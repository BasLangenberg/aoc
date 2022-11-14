package gameconsole

import (
	"fmt"
	"strconv"
	"strings"
)

// Console models a handheld game console
type Console struct {
	Program          []string
	Pointer          int
	Previouspointers []int
	Acu              int
}

// New returns a pointer to a Console instance
func New(program []string) *Console {
	return &Console{
		Program: program,
		Pointer: 0,
		Acu:     0,
	}
}

// Pulse executes a clock cycle in the console
func (c *Console) Pulse() bool {
	if c.Pointer > len(c.Program)-1 {
		fmt.Printf("Program done, ACU: %v\n", c.Acu)
		return true
	}

	instructions := strings.Split(c.Program[c.Pointer], " ")
	switch instructions[0] {
	case "nop":
		c.Pointer++
	case "acc":
		num, err := strconv.Atoi(instructions[1][1:])
		if instructions[1][0] == '+' {
			if err != nil {
				fmt.Printf("Unable to parse integer: %v", err)
				c.Pointer++
			}
			c.Acu += num
			c.Previouspointers = append(c.Previouspointers, c.Pointer)
			c.Pointer++
		}
		if instructions[1][0] == '-' {
			if err != nil {
				fmt.Printf("Unable to parse integer: %v", err)
				c.Pointer++
			}
			c.Acu -= num
			c.Previouspointers = append(c.Previouspointers, c.Pointer)
			c.Pointer++
		}
		if instructions[1][0] != '-' && instructions[1][0] != '+' {
			fmt.Printf("Unable to parse ACU instruction: %v", instructions[1][0])
			c.Previouspointers = append(c.Previouspointers, c.Pointer)
			c.Pointer++
		}
	case "jmp":
		num, err := strconv.Atoi(instructions[1][1:])
		if instructions[1][0] == '+' {
			if err != nil {
				fmt.Printf("Unable to parse integer: %v", err)
				c.Previouspointers = append(c.Previouspointers, c.Pointer)
				c.Pointer++
			}
			c.Previouspointers = append(c.Previouspointers, c.Pointer)
			c.Pointer += num
		}
		if instructions[1][0] == '-' {
			if err != nil {
				fmt.Printf("Unable to parse integer: %v", err)
				c.Pointer++
			}
			c.Pointer -= num
		}
		if instructions[1][0] != '-' && instructions[1][0] != '+' {
			fmt.Printf("Unable to parse JMP instruction: %v", instructions[1][0])
			c.Previouspointers = append(c.Previouspointers, c.Pointer)
			c.Pointer++
		}
	default:
		fmt.Printf("Unable to parse instruction: %v", instructions[0])
		c.Previouspointers = append(c.Previouspointers, c.Pointer)
		c.Pointer++
	}

	return false

}
