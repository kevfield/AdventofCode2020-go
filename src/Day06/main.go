package main

import "fmt"

func main() {

	// get input from file
	puzzleInput, _ := readFile("input.txt")

	// run part1
	p1answer := processLandingCards(puzzleInput)
	fmt.Println("Part 1 =", p1answer)

}
