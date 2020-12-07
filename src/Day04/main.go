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
		getpassportData(puzzleInput)

	} else if partid == "b" {

	} else {
		fmt.Println("Invalid part id, please specify eiter a or b")
	}

}
