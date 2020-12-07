package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// read a file from an input and return into a slice of strings
func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// use flags to get user input
func inputFlags() (string, string) {

	// variables declaration
	var flagfile string
	var partid string

	// flags declaration using flag package
	flag.StringVar(&flagfile, "file", "input.txt", "filename of the input data eg: input.txt")
	flag.StringVar(&partid, "part", "a", "use either a or b")

	flag.Parse() // after declaring flags we need to call it
	return flagfile, partid
}

func getpassportData(passports []string) {

	passportNumber := 0
	var p [1000]passportList
	validItems := 0
	validPassports := 0
OuterLoop:
	for _, line := range passports {

		if len(line) == 0 {
			passportNumber++
		} else {
			for _, i := range strings.SplitAfter(line, " ") {
				fmt.Sscanf(i, "byr:%s", &p[passportNumber].byr)
				fmt.Sscanf(i, "iyr:%s", &p[passportNumber].iyr)
				fmt.Sscanf(i, "eyr:%s", &p[passportNumber].eyr)
				fmt.Sscanf(i, "pid:%s", &p[passportNumber].pid)
				fmt.Sscanf(i, "cid:%s", &p[passportNumber].cid)
				fmt.Sscanf(i, "hgt:%s", &p[passportNumber].hgt)
				fmt.Sscanf(i, "hcl:%s", &p[passportNumber].hcl)
				fmt.Sscanf(i, "ecl:%s", &p[passportNumber].ecl)
				if len(p[passportNumber].byr) != 0 {
					validItems++
				}
				fmt.Println(p[passportNumber].byr)
				break OuterLoop
				if len(p[passportNumber].iyr) != 0 {
					validItems++
				}
				if len(p[passportNumber].eyr) != 0 {
					validItems++
				}
				if len(p[passportNumber].pid) != 0 {
					validItems++
				}
				//if len(p.cid) != 0 {
				//	validItems++
				//}
				if len(p[passportNumber].hgt) != 0 {
					validItems++
				}
				if len(p[passportNumber].hcl) != 0 {
					validItems++
				}
				if len(p[passportNumber].ecl) != 0 {
					validItems++
				}
				if validItems == 7 {
					validPassports++
				}
				break OuterLoop
			}

		}

	}
	fmt.Println(validPassports)
}
