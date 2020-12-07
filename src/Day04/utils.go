package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func getpassportData(passports []string, part string) (int, int) {

	passportNumber := 0
	var p [1000]passportList
	pavalidItems := 0
	pbvalidItems := 0
	partaValidPassports := 0
	partbValidPassports := 0
	charMatch := 0

	for _, line := range passports {
		if len(line) == 0 {
			passportNumber++
		} else {
			for _, i := range strings.Split(line, " ") {
				//fmt.Println("i =", i)
				fmt.Sscanf(i, "byr:%s", &p[passportNumber].byr)
				fmt.Sscanf(i, "iyr:%s", &p[passportNumber].iyr)
				fmt.Sscanf(i, "eyr:%s", &p[passportNumber].eyr)
				fmt.Sscanf(i, "pid:%s", &p[passportNumber].pid)
				fmt.Sscanf(i, "cid:%s", &p[passportNumber].cid)
				fmt.Sscanf(i, "hgt:%s", &p[passportNumber].hgt)
				fmt.Sscanf(i, "hcl:%s", &p[passportNumber].hcl)
				fmt.Sscanf(i, "ecl:%s", &p[passportNumber].ecl)
			}

		}

	}

	for j := 0; j < len(p); j++ {
		pavalidItems = 0
		pbvalidItems = 0
		charMatch = 0

		// part a
		if len(p[j].byr) != 0 {
			pavalidItems++
		}
		// part b
		intByr, _ := strconv.Atoi(p[j].byr)
		if len(p[j].byr) == 4 && intByr >= 1920 && intByr <= 2002 {
			pbvalidItems++
		}

		// part a
		if len(p[j].iyr) != 0 {
			pavalidItems++
		}
		// part b
		intIyr, _ := strconv.Atoi(p[j].iyr)
		if len(p[j].iyr) == 4 && intIyr >= 2010 && intIyr <= 2020 {
			pbvalidItems++
		}

		// part a
		if len(p[j].eyr) != 0 {
			pavalidItems++
		}
		// part b
		intEyr, _ := strconv.Atoi(p[j].eyr)
		if len(p[j].eyr) == 4 && intEyr >= 2020 && intEyr <= 2030 {
			pbvalidItems++
		}

		// part a
		if len(p[j].pid) != 0 {
			pavalidItems++
		}
		// part b
		if len(p[j].pid) == 9 {
			if _, err := strconv.Atoi(p[j].pid); err == nil {
				pbvalidItems++
			}
		}

		//if len(p.cid) != 0 {
		//	validItems++
		//}

		//part a
		if len(p[j].hgt) != 0 {
			pavalidItems++
		}
		// part b
		var getHeight int
		if strings.Index(p[j].hgt, "in") > 0 {
			fmt.Sscanf(p[j].hgt, "%din", &getHeight)
			if getHeight >= 59 && getHeight <= 76 {
				pbvalidItems++
			}
		}

		if strings.Index(p[j].hgt, "cm") > 0 {
			fmt.Sscanf(p[j].hgt, "%dcm", &getHeight)
			if getHeight >= 150 && getHeight <= 193 {
				pbvalidItems++
			}
		}

		// part a
		if len(p[j].hcl) != 0 {
			pavalidItems++
		}
		// part b
		if len(p[j].hcl) != 0 {
			if p[j].hcl[0] == 35 {
				if len(p[j].hcl) == 7 {
					re := regexp.MustCompile(`[a-f0-9]`)
					for l := 1; l <= len(p[j].hcl)-1; l++ {
						if re.MatchString(p[j].hcl) == true {
							charMatch++
						}
					}
				}
			}

			if charMatch == 6 {
				pbvalidItems++
			}
		}

		// part a
		if len(p[j].ecl) != 0 {
			pavalidItems++
		}
		// part b
		if p[j].ecl == "amb" || p[j].ecl == "blu" || p[j].ecl == "brn" || p[j].ecl == "gry" || p[j].ecl == "grn" || p[j].ecl == "hzl" || p[j].ecl == "oth" {
			pbvalidItems++
		}

		/////
		if pavalidItems == 7 {
			partaValidPassports++
		}
		if pbvalidItems == 7 {
			partbValidPassports++
		}
	}
	return partaValidPassports, partbValidPassports
}
