package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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
	//lines = append(lines, "")
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

func processLandingCards(input []string) {

	var groupsCards map[byte]int
	var countResponses int
	var totalResponses int

	groupsCards = make(map[byte]int)

	for _, person := range input {
		if person == "" {
			countResponses = 0
			for _, value := range groupsCards {
				if value >= 1 {
					countResponses++
				}
			}
			totalResponses = totalResponses + countResponses
			// found new group
			// make new map
			groupsCards = make(map[byte]int)
			// reset count responses

		} else {
			//countResponses++
		}
		// loop through lines to get answers and add to map
		for _, answer := range person {
			groupsCards[byte(answer)]++
		}

	}
	// edge case, count last line if not counted
	countResponses = 0
	for _, value := range groupsCards {
		if value >= 1 {
			countResponses++
		}
	}
	totalResponses = totalResponses + countResponses

	fmt.Println(totalResponses)
}
