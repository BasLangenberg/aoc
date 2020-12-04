package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type passport struct {
	byr        string
	iyr        string
	eyr        string
	hgt        string
	hcl        string
	ecl        string
	pid        string
	cid        string
	partAValid bool
	partBValid bool
}

func main() {
	passports := []passport{}
	input, err := ioutil.ReadFile("input")

	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}

	sinput := string(input)
	pssprts := strings.Split(sinput, "\n\n")
	fmt.Printf("total passports: %v\n", len(pssprts))

	for _, pssprt := range pssprts {
		pssprt = strings.ReplaceAll(pssprt, "\n", " ")
		var pp passport
		for _, item := range strings.Split(pssprt, " ") {
			switch strings.Split(item, ":")[0] {
			case "byr":
				pp.byr = strings.Split(item, ":")[1]
			case "iyr":
				pp.iyr = strings.Split(item, ":")[1]
			case "eyr":
				pp.eyr = strings.Split(item, ":")[1]
			case "hgt":
				pp.hgt = strings.Split(item, ":")[1]
			case "hcl":
				pp.hcl = strings.Split(item, ":")[1]
			case "ecl":
				pp.ecl = strings.Split(item, ":")[1]
			case "pid":
				pp.pid = strings.Split(item, ":")[1]
			case "cid":
				pp.cid = strings.Split(item, ":")[1]
			}
		}

		pp.validatePartA()
		pp.validatePartB()
		fmt.Printf("%+v\n", pp)
		passports = append(passports, pp)
	}

	// Solve part 1
	var valid int
	var validb int
	for _, pssprt := range passports {
		if pssprt.partAValid {
			valid++
		}
	}

	for _, pssprt := range passports {
		if pssprt.partBValid {
			validb++
		}
	}

	fmt.Printf("Valid passports part 1: %v\n", valid)
	fmt.Printf("Valid passports part 2: %v\n", validb)
}

func (pw *passport) validatePartA() {
	if pw.byr == "" || pw.iyr == "" || pw.eyr == "" || pw.hgt == "" || pw.hcl == "" || pw.ecl == "" || pw.pid == "" {
		pw.partAValid = false
		return
	}
	pw.partAValid = true
}

func (pw *passport) validatePartB() {

	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	byr, err := strconv.Atoi(pw.byr)
	if err != nil {
		return
	}

	if byr < 1920 || byr > 2002 {
		return
	}

	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	iyr, err := strconv.Atoi(pw.iyr)
	if err != nil {
		return
	}

	if iyr < 2010 || iyr > 2020 {
		return
	}

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	eyr, err := strconv.Atoi(pw.eyr)
	if err != nil {
		return
	}

	if eyr < 2020 || eyr > 2030 {
		return
	}

	// hgt (Height) - a number followed by either cm or in:
	//     If cm, the number must be at least 150 and at most 193.
	//     If in, the number must be at least 59 and at most 76.
	if len(pw.hgt) == 0 {
		return
	}

	indicator := pw.hgt[len(pw.hgt)-2:]
	if indicator == "cm" {
		num, err := strconv.Atoi(pw.hgt[:len(pw.hgt)-2])
		if err != nil {
			return
		}
		if num < 150 || num > 193 {
			return
		}
	}

	if indicator == "in" {
		num, err := strconv.Atoi(pw.hgt[:len(pw.hgt)-2])
		if err != nil {
			return
		}
		if num < 59 || num > 76 {
			return
		}
	}

	if indicator != "in" && indicator != "cm" {
		return
	}

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	if len(pw.hcl) == 0 {
		return
	}

	if pw.hcl[0] != byte('#') {
		return
	}
	for _, charact := range pw.hcl[1:] {
		switch charact {
		case 'a', 'b', 'c', 'd', 'e', 'f', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			continue
		default:
			return
		}
	}
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	switch pw.ecl {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		break
	default:
		return
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	if len(pw.pid) != 9 {
		return
	}

	for _, num := range pw.pid {
		n, err := strconv.Atoi(string(num))
		if err != nil {
			return
		}

		switch n {
		case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
			break
		default:
			return
		}
	}
	// cid (Country ID) - ignored, missing or not.

	// All ok?
	pw.partBValid = true

}
