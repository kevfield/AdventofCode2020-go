package main

import (
	"fmt"
)

type passportSlice []string

type passportList struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func main() {

	// get user input
	inputLocation, partid := inputFlags()

	// get puzzle data file
	puzzleInput, _ := readFile(inputLocation)

	if partid == "a" {
		// get the data from the input file
		p1Result, _ := getpassportData(puzzleInput, partid)
		fmt.Println("Result =", p1Result)

	} else if partid == "b" {
		_, p2result := getpassportData(puzzleInput, partid)
		fmt.Println("Result = ", p2result)

	} else {
		fmt.Println("Invalid part id, please specify eiter a or b")
	}

}
