package main

import (
	"fmt"
	"strconv"
)

func convertInputtoInt(strInput []string) []int {
	var intSlice []int
	for a := 0; a < len(strInput); a++ {
		b, _ := strconv.Atoi(strInput[a])
		intSlice = append(intSlice, b)
	}
	return intSlice
}

func findNumber(puzInput []int) int {

	preambleLen := 25
	preamstartingPos := 0
	preamendingPos := preamstartingPos + preambleLen
	foundAnswer := false
	validCount := 0
	foundResult := 0

	for foundAnswer == false {
		for i := preamstartingPos; i <= preamendingPos; i++ {
			for j := preamstartingPos; j <= preamendingPos; j++ {
				if puzInput[i] != puzInput[j] {
					if puzInput[i]+puzInput[j] == puzInput[preamendingPos+1] {
						// found valid number, move to next set
						validCount++
					}
				}
			}
		}
		if validCount == 0 {
			foundAnswer = true
			foundResult = puzInput[preamendingPos+1]
		}
		// increment starting pos
		preamstartingPos++
		// reset valid count
		validCount = 0
		// set new ending
		preamendingPos = preamstartingPos + preambleLen
	}
	return foundResult
}

func contiguousNumbers(p1Answer int, p2Input []int) int {
	p2foundResult := 0
	startingPos := 0
	countinSection := 0
	endingPos := 0
	runningTotal := 0
	howmanyRuns := 0
	matchedP1 := false

	for matchedP1 == false {
		for c := startingPos; c <= len(p2Input); c++ {
			// update running total
			runningTotal = runningTotal + p2Input[c]
			// Count how many numbers added in this run
			countinSection++
			if runningTotal > p1Answer {
				// exceeded target amount, change starting number and start again
				runningTotal = 0
				break
			} else if runningTotal == p1Answer {
				// target amount found, find starting and ending pos of run
				startingPos = c - countinSection
				endingPos = c
				matchedP1 = true
				break
			}
		}
		howmanyRuns++
		startingPos++
		countinSection = 0
	}

	// find lowest number between the two positions
	smallest := p2Input[startingPos]
	largest := p2Input[startingPos]
	// find smallest and largest numbers out of section
	for d := startingPos; d <= endingPos; d++ {
		if p2Input[d] > smallest {
			smallest = p2Input[d]
		} else if p2Input[d] < largest {
			largest = p2Input[d]
		}
	}
	p2foundResult = smallest + largest
	return p2foundResult
}

func main() {

	// get input from user
	fileLoc, _ := inputFlags()

	// get puzzle input
	puzzleinput, _ := readFile(fileLoc)

	// convert puzzle input to slice of Int
	intInput := convertInputtoInt(puzzleinput)

	// Part 1
	p1puzAnswer := findNumber(intInput)
	fmt.Println("Part 1: ", p1puzAnswer)

	// Part 2
	fmt.Println("Part 2: ", contiguousNumbers(p1puzAnswer, intInput))
}
