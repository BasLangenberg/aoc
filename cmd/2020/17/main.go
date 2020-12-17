package main

import (
	"fmt"
	"math"
)

type conwaycube struct {
	x, y, z int
}

type ccubes map[conwaycube]bool

type cubedynamic struct {
	xmin, xmax, ymin, ymax, zmin, zmax int
}

func main() {
	cubes := make(ccubes)
	// input := []string{
	// 	"##......",
	// 	".##...#.",
	// 	".#######",
	// 	"..###.##",
	// 	".#.###..",
	// 	"..#.####",
	// 	"##.####.",
	// 	"##..#.##",
	// }
	input := []string{
		".#.",
		"..#",
		"###",
	}

	for pos, str := range input {
		for ps, charact := range str {
			var active bool

			if charact == '.' {
				active = false
			}

			if charact == '#' {
				active = true
			}

			cc := conwaycube{
				z: 0,
				y: pos,
				x: ps,
			}

			mincc := conwaycube{
				z: -1,
				y: pos,
				x: ps,
			}

			pluscc := conwaycube{
				z: 1,
				y: pos,
				x: ps,
			}

			cubes[cc] = active
			cubes[mincc] = false
			cubes[pluscc] = false
		}
	}

	for i := 0; i < 6; i++ {
		bound := getMinMax(cubes)
		fmt.Println(bound)
		pulse(cubes, bound)
		printGrid(cubes, bound)
	}

	var count int

	for key, val := range cubes {
		if val {
			fmt.Printf("%+v - %v\n", key, val)
			count++
		}
	}

	// fmt.Println("Too high: 804")
	fmt.Printf("Part A: %v\n", count)

}

func printGrid(c ccubes, b cubedynamic) {
	for z := b.zmin; z <= b.zmax; z++ {
		fmt.Printf("x: %v", z)
		for y := b.ymin; y <= b.ymax; y++ {

			for x := b.xmin; x <= b.xmax; x++ {
				cc := conwaycube{
					z: z,
					x: x,
					y: y,
				}

				val, ok := c[cc]
				if !ok {
					fmt.Printf(".")
					break
				}

				if val {
					fmt.Printf("#")
					break
				}

				fmt.Printf(".")
			}

			fmt.Printf("\n")

		}
		fmt.Println()
		fmt.Println()
		fmt.Println()
	}
}

func getMinMax(c ccubes) cubedynamic {
	zmax := math.MinInt64
	ymax := math.MinInt64
	xmax := math.MinInt64

	zmin := math.MaxInt64
	ymin := math.MaxInt64
	xmin := math.MaxInt64

	for key := range c {
		if zmax < key.z {
			zmax = key.z
		}
		if ymax < key.y {
			ymax = key.y
		}
		if xmax < key.x {
			xmax = key.x
		}
		if zmin > key.z {
			zmin = key.z
		}
		if ymin > key.y {
			ymin = key.y
		}
		if xmin > key.x {
			xmin = key.x
		}
	}

	return cubedynamic{
		xmin: xmin,
		xmax: xmax,
		ymin: ymin,
		ymax: ymax,
		zmin: zmin,
		zmax: zmax,
	}
}

func pulse(c ccubes, b cubedynamic) {
	updates := make(map[conwaycube]bool)

	for key, val := range c {
		var active int
		for z := key.z - 1; z <= key.z+1; z++ {
			for y := key.y - 1; y <= key.y+1; y++ {
				for x := key.x - 1; x <= key.x+1; x++ {
					cc := conwaycube{
						x: x,
						y: y,
						z: z,
					}

					if cc == key {
						break
					}

					val, ok := c[cc]
					if !ok {
						updates[cc] = false
						break
					}

					if val == true {
						active++
					}
				}
			}
		}

		if !(active == 3 || active == 2) && val {
			updates[key] = false
		}

		if !val && active == 3 {
			updates[key] = true
		}
	}

	for key, val := range updates {
		c[key] = val
	}
}
