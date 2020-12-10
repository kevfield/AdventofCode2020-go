package main

import "fmt"

func processOperations(pInput []string) {

	usedInstructions := make([]bool, len(pInput))
	//startingPoint := usedInstructions[0]
	currentPos := 0
	accumulator := 0

	var accAmount, jmpAmount int
	var foundInstruction string

	for usedInstructions[currentPos] == false {
		fmt.Sscanf(pInput[currentPos], "%s", &foundInstruction)

		if foundInstruction == "acc" {
			// get acc instruction
			fmt.Sscanf(pInput[currentPos], "acc%d", &accAmount)
			usedInstructions[currentPos] = true
			accumulator = accumulator + accAmount
			currentPos = currentPos + 1

		} else if foundInstruction == "nop" {
			usedInstructions[currentPos] = true
			currentPos = currentPos + 1

		} else if foundInstruction == "jmp" {
			fmt.Sscanf(pInput[currentPos], "jmp%d", &jmpAmount)
			usedInstructions[currentPos] = true
			currentPos = currentPos + jmpAmount

		}

		// reset variables
		foundInstruction = ""
		accAmount = 0
		jmpAmount = 0

	}
	fmt.Println(accumulator)
}

func main() {

	// Grab user choices
	inputFile, _ := inputFlags()
	// pull in puzzle input
	puzzleInput, _ := readFile(inputFile)

	processOperations(puzzleInput)

}
