package main

import (
	"bufio"
	"flag"
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

func processLandingCards(input []string) (int, int) {

	var groupsCards map[byte]int
	var p1countResponses int
	var p1totalResponses int
	var p2countResponses int
	var p2totalResponses int
	newPerson := 0

	groupsCards = make(map[byte]int)

	for _, person := range input {
		if person == "" {
			// reset count responses
			p1countResponses = 0
			p2countResponses = 0
			for _, value := range groupsCards {
				if value >= 1 {
					p1countResponses++
				}
			}
			for _, value := range groupsCards {
				if value >= newPerson {
					p2countResponses++
				}
			}
			newPerson = 0
			p1totalResponses = p1totalResponses + p1countResponses
			p2totalResponses = p2totalResponses + p2countResponses
			// found new group
			// make new map
			groupsCards = make(map[byte]int)
		} else {
			newPerson++
		}
		// loop through lines to get answers and add to map
		for _, answer := range person {
			groupsCards[byte(answer)]++
		}

	}
	// edge case p1, count last line
	p1countResponses = 0
	for _, value := range groupsCards {
		if value >= 1 {
			p1countResponses++
		}
	}
	p1totalResponses = p1totalResponses + p1countResponses

	// edge p2
	p2countResponses = 0
	for _, value := range groupsCards {
		if value >= newPerson {
			p2countResponses++
		}
	}
	p2totalResponses = p2totalResponses + p2countResponses
	return p1totalResponses, p2totalResponses
}
