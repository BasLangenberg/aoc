package main

import "fmt"

type conwaycube struct {
	x, y, z int
}

type ccubes map[conwaycube]bool

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
			cubes[mincc] = active
			cubes[pluscc] = active
		}
	}

	for i := 0; i < 6; i++ {
		pulse(cubes)
	}

	var count int

	for _, val := range cubes {
		if val {
			count++
		}
	}

	// fmt.Println("Too high: 804")
	fmt.Printf("Part A: %v\n", count)

}

// TODO: Loop through coordinates instead of through map!
func pulse(c ccubes) {
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
						c[cc] = false
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
