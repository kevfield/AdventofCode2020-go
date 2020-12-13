package main

import "fmt"

func p1findJolts(p1Input []int) int {
	joltDiffs := 0
	oneCount := 0
	threeCount := 0

	myAdapter := p1Input[len(p1Input)-1]
	myValue := myAdapter + 3
	p1Input = append(p1Input, myValue)

	for i := 0; i < len(p1Input)-1; i++ {

		j := i + 1
		if p1Input[j]-p1Input[i] == 1 {
			oneCount++
		} else if p1Input[j]-p1Input[i] == 3 {
			threeCount++
		}
	}
	// grab first one
	if p1Input[1]-p1Input[0] == 1 {
		oneCount++
	} else {
		threeCount++
	}
	joltDiffs = oneCount * threeCount
	return joltDiffs
}

func main() {

	// get input from user
	fileLoc, _ := inputFlags()

	// get puzzle input
	puzzleinput, _ := readFile(fileLoc)

	// convert puzzle input to slice of Int
	intInput := convertInputtoInt(puzzleinput)

	fmt.Println("Part 1: ", p1findJolts(intInput))

}
