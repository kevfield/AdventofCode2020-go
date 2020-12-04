package main

import "fmt"

func main() {

	// get user input
	filelocation, partchoice, travelX, travelY := inputFlags()

	// get puzzle input from file
	puzzleInput, _ := readFile(filelocation)

	// get number of trees
	if partchoice == "a" {

		part1Answer := treeCounter(puzzleInput, travelX, travelY)
		fmt.Println("Part a: Number of Trees: ", part1Answer)

	} else if partchoice == "b" {

		part2Answer := treeCounter(puzzleInput, 1, 1)
		part2Answer *= treeCounter(puzzleInput, 3, 1)
		part2Answer *= treeCounter(puzzleInput, 5, 1)
		part2Answer *= treeCounter(puzzleInput, 7, 1)
		part2Answer *= treeCounter(puzzleInput, 1, 2)

		fmt.Println("Part b: Number of Trees", part2Answer)

	} else {
		fmt.Println("Invalid part choice, please specify a or b")
	}

}
