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
	return lines, scanner.Err()
}

// use flags to get user input
func inputFlags() (string, string, int, int) {

	// variables declaration
	var flagfile string
	var partid string
	var travelx int
	var travely int

	// flags declaration using flag package
	flag.StringVar(&flagfile, "file", "input.txt", "filename of the input data eg: input.txt")
	flag.StringVar(&partid, "part", "a", "use either a or b")
	flag.IntVar(&travelx, "travelx", 3, "specify distance to travel to the right")
	flag.IntVar(&travely, "travely", 1, "specify distance to travel down")

	flag.Parse() // after declaring flags we need to call it
	return flagfile, partid, travelx, travely
}

// find how many trees
func treeCounter(mapofTrees []string, xDistance int, yDistance int) int {

	maxX := len(mapofTrees[0])
	maxY := len(mapofTrees)

	currentXPos := xDistance
	currentYPos := yDistance
	keepWalking := true
	numberofTrees := 0

	for keepWalking == true {
		if currentYPos >= maxY {
			keepWalking = false
		} else {
			if mapofTrees[currentYPos][currentXPos] == '#' {
				numberofTrees++
			}
		}
		currentXPos = (currentXPos + xDistance) % maxX
		currentYPos = currentYPos + yDistance

	}
	return numberofTrees
}
