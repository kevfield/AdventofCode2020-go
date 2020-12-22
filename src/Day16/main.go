package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ticketfieldRanges struct {
	lowminVal  int
	lowmaxVal  int
	highminVal int
	highmaxVal int
}

func validateTickets(ticketInput []string) {

	var grabstuff, grabstuff2, grabstuff3, grabstuff4 int
	var myTicket int
	var nearbyTickets int
	var structStuff []ticketfieldRanges
	var invalidFields []int

	// loop through file and grab starting points
	for i := 0; i < len(ticketInput); i++ {
		if ticketInput[i] == "your ticket:" {
			// my ticket location found
			myTicket = i + 1
		} else if ticketInput[i] == "nearby tickets:" {
			// all other tickets starting point found
			nearbyTickets = i + 1
		}
	}
	// delclare structs
	structStuff = make([]ticketfieldRanges, myTicket-2)

	// setup struct data and loop through nearby tickets
	for line := 0; line < myTicket-2; line++ {

		// split input to get values
		splitString := strings.Split(ticketInput[line], ":")
		fmt.Sscanf(splitString[1], "%d-%d or %d-%d", &grabstuff, &grabstuff2, &grabstuff3, &grabstuff4)
		// add values to slice of structs
		structStuff[line].lowminVal = grabstuff
		structStuff[line].lowmaxVal = grabstuff2
		structStuff[line].highminVal = grabstuff3
		structStuff[line].highmaxVal = grabstuff4
	}

	// loop through nearby tickets to test for invalid
	continueLoop := true
	var nearbySplit []string
	stringPos := 0
	runCount := 0

	for continueLoop == true {
		for j := nearbyTickets; j < len(ticketInput); j++ {
			nearbySplit = strings.Split(ticketInput[j], ",")
			nearbySplitint, _ := strconv.Atoi(nearbySplit[stringPos])

			if nearbySplitint >= structStuff[runCount].lowminVal && nearbySplitint <= structStuff[runCount].lowmaxVal || nearbySplitint >= structStuff[runCount].highminVal && nearbySplitint <= structStuff[runCount].highmaxVal {
				// value is valid, do nothing
			} else {
				// value not valid, add to slice for later
				invalidFields = append(invalidFields, nearbySplitint)

			}
			//fmt.Println(structStuff[runCount])
			//fmt.Println("sringpos", stringPos)
			//fmt.Println("j", j)

			// check if we're at the end of the file
			if stringPos == len(nearbySplit)-1 && j == len(ticketInput)-1 {
				continueLoop = false
				//fmt.Println("hi")
				break
			}
		}
		stringPos++
		runCount++

	}

	// process invalids
	fmt.Println("invalid struct", invalidFields)
	invalids := 0
	for l := 0; l < len(invalidFields); l++ {
		invalids = invalidFields[l] + invalids
	}
	fmt.Println("Part 1:", invalids)

}

func main() {

	// get input from user
	fileLoc, _ := inputFlags()
	// get puzzle input
	puzzleinput, _ := readFile(fileLoc)

	validateTickets(puzzleinput)

}
