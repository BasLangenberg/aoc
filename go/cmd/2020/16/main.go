package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/BasLangenberg/aoc/internal/dataconversion"
)

type boundary struct {
	label   string
	up, low []int
	pos     int
}

type ticket struct {
	tic []int
	ok  bool
}

func main() {
	input, err := ioutil.ReadFile("input")

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
		os.Exit(1)
	}

	data := strings.Split(string(input), "\n\n")
	boundaries := []boundary{}
	for _, bound := range strings.Split(data[0], "\n") {
		bound = strings.Replace(bound, " or ", ",", -1)
		boundpair := strings.Split(bound, ": ")

		x := boundary{
			pos:   -1,
			label: boundpair[0],
		}

		for _, b := range strings.Split(boundpair[1], ",") {
			t := strings.Split(b, "-")

			t1, err := strconv.Atoi(t[0])
			if err != nil {
				fmt.Printf("Unable to parse int: %v", err)
				os.Exit(1)
			}

			t2, err := strconv.Atoi(t[1])
			if err != nil {
				fmt.Printf("Unable to parse int: %v", err)
				os.Exit(1)
			}

			x.low = append(x.low, t1)
			x.up = append(x.up, t2)
		}

		boundaries = append(boundaries, x)
	}

	var numGoodTicks int

	// Part A
	var sum int
	var Ticks []ticket
	for pos, tick := range strings.Split(data[2], "\n") {
		if pos == 0 {
			continue
		}

		var nums []int

		for _, num := range strings.Split(tick, ",") {
			n, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("Unable to parse int: %v", err)
				os.Exit(1)
			}

			nums = append(nums, n)
		}

		t := ticket{
			tic: nums,
			ok:  true,
		}

		for _, num := range nums {
			var ok bool
			for _, bnd := range boundaries {
				for i := 0; i < len(bnd.up); i++ {
					if num >= bnd.low[i] && num <= bnd.up[i] {
						ok = true
					}
				}
			}

			if !ok {
				sum += num
				t.ok = false
				continue
			}
		}

		if t.ok {
			numGoodTicks++
		}

		Ticks = append(Ticks, t)

	}

	fmt.Printf("Part A: %+v\n", sum)

	// Part B
	type mark struct {
		num int
		ok  bool
	}

	type positions struct {
		num     int
		columns []int
	}

	possiblePositions := []positions{}

	for pos, bnd := range boundaries {
		var nums []int
		for i := 0; i < len(Ticks[0].tic); i++ {
			marks := []mark{}
			for _, tick := range Ticks {
				if !tick.ok {
					continue
				}

				for y := 0; y < len(bnd.up); y++ {
					if tick.tic[i] >= bnd.low[y] && tick.tic[i] <= bnd.up[y] {
						marks = append(marks, mark{tick.tic[i], true})
						break
					}
				}
			}

			if len(marks) == numGoodTicks {
				nums = append(nums, i)
			}
		}
		possiblePositions = append(possiblePositions, positions{
			num:     pos,
			columns: nums,
		})
	}

	foundNums := []int{math.MaxInt64}
	for i := 0; i < 20; i++ {
		for _, pp := range possiblePositions {
			if len(pp.columns)-i == 1 {
				var fnumb int
				for _, nm := range pp.columns {
					if !dataconversion.IntInSlice(nm, foundNums) {
						fnumb = nm
						break
					}
				}
				boundaries[pp.num].pos = fnumb
				foundNums = append(foundNums, fnumb)
			}
		}
	}

	sumB := 1
	myTicket := strings.Split(data[1], "\n")[1]
	myTicketNums := dataconversion.StringToIntArray(myTicket)

	fmt.Println(len(boundaries))
	for _, bnd := range boundaries {
		fmt.Printf("%+v\n", bnd)
	}

	for _, bnd := range boundaries {
		if bnd.label[:3] == "dep" {
			sumB *= myTicketNums[bnd.pos]
		}
	}

	fmt.Printf("Part B: %v", sumB)

}
