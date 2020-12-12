package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type instruction struct {
	Dir   string
	Steps int
}

type ferry struct {
	Facing  string
	Pos     []int
	PrevPos [][]int
}

func main() {
	input, err := os.Open("input")
	defer input.Close()

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var instr []instruction

	ship := ferry{
		Facing: "E",
		Pos:    []int{0, 0},
	}

	for scanner.Scan() {
		line := scanner.Text()
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Printf("Unable to parse line: %v", err)
			os.Exit(1)
		}

		inst := instruction{
			Dir:   string(line[0]),
			Steps: steps,
		}

		instr = append(instr, inst)

	}

	for _, inst := range instr {
		ship.PrevPos = append(ship.PrevPos, []int{ship.Pos[0], ship.Pos[1]})
		ship.move(&inst)
	}

	ship.PrevPos = append(ship.PrevPos, []int{ship.Pos[0], ship.Pos[1]})
	// fmt.Printf("%+v\n", ship.PrevPos)

	fmt.Printf("Manhattan distance: %v, %v\n", ship.Pos[0], ship.Pos[1])
	fmt.Println("Too low: 1341")
	fmt.Println("Too low: 1465")
	fmt.Printf("Sum Part A: %v\n", abs(ship.Pos[0])+abs(ship.Pos[1]))

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (s *ferry) move(inst *instruction) {
	// Action N means to move north by the given value.
	// Action S means to move south by the given value.
	// Action E means to move east by the given value.
	// Action W means to move west by the given value.
	// Action L means to turn left the given number of degrees.
	// Action R means to turn right the given number of degrees.
	// Action F means to move forward by the given value in the direction the ship is currently facing.

	switch inst.Dir {
	case "N":
		s.Pos[0] -= inst.Steps
	case "S":
		s.Pos[0] += inst.Steps
	case "E":
		s.Pos[1] += inst.Steps
	case "W":
		s.Pos[1] -= inst.Steps
	case "F":
		switch s.Facing {
		case "N":
			s.Pos[0] -= inst.Steps
		case "S":
			s.Pos[0] += inst.Steps
		case "E":
			s.Pos[1] += inst.Steps
		case "W":
			s.Pos[1] -= inst.Steps
		default:
			fmt.Printf("Unable to parse instruction %v - %v", inst.Dir, inst.Steps)
		}
	case "R":
		switch s.Facing {
		case "N":
			switch inst.Steps {
			case 90:
				s.Facing = "E"
			case 180:
				s.Facing = "S"
			case 270:
				s.Facing = "W"
			default:
				fmt.Printf("Unable to parse instruction %v - %v", inst.Dir, inst.Steps)
			}
		case "E":
			switch inst.Steps {
			case 90:
				s.Facing = "S"
			case 180:
				s.Facing = "W"
			case 270:
				s.Facing = "N"
			default:
				fmt.Printf("Unable to parse instruction %v - %v", inst.Dir, inst.Steps)
			}
		case "S":
			switch inst.Steps {
			case 90:
				s.Facing = "W"
			case 180:
				s.Facing = "N"
			case 270:
				s.Facing = "E"
			default:
				fmt.Printf("Unable to parse instruction %v - %v", inst.Dir, inst.Steps)
			}
		case "W":
			switch inst.Steps {
			case 90:
				s.Facing = "N"
			case 180:
				s.Facing = "E"
			case 270:
				s.Facing = "S"
			default:
				fmt.Printf("Unable to parse instruction %v - %v", inst.Dir, inst.Steps)
			}
		}
	case "L":
		switch s.Facing {
		case "N":
			switch inst.Steps {
			case 90:
				s.Facing = "W"
			case 180:
				s.Facing = "S"
			case 270:
				s.Facing = "E"
			default:
				fmt.Printf("Unable to parse instruction %v - %v", inst.Dir, inst.Steps)
			}
		case "E":
			switch inst.Steps {
			case 90:
				s.Facing = "N"
			case 180:
				s.Facing = "W"
			case 270:
				s.Facing = "S"
			default:
				fmt.Printf("Unable to parse instruction %v - %v", inst.Dir, inst.Steps)
			}
		case "S":
			switch inst.Steps {
			case 90:
				s.Facing = "E"
			case 180:
				s.Facing = "N"
			case 270:
				s.Facing = "W"
			default:
				fmt.Printf("Unable to parse instruction %v - %v", inst.Dir, inst.Steps)
			}
		case "W":
			switch inst.Steps {
			case 90:
				s.Facing = "S"
			case 180:
				s.Facing = "E"
			case 270:
				s.Facing = "W"
			default:
				fmt.Printf("Unable to parse instruction %v - %v", inst.Dir, inst.Steps)
			}
		default:
			fmt.Printf("Unable to parse instruction %v", inst.Dir)
		}
	}
}
