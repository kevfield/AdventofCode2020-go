package main

import "fmt"

func processOperations(pInput []string) (int, bool) {

	usedInstructions := make([]bool, len(pInput))
	//startingPoint := usedInstructions[0]
	currentPos := 0
	p1accumulator := 0
	var p2Found bool

	var accAmount, jmpAmount int
	var foundInstruction string

	for usedInstructions[currentPos] == false {
		if currentPos > len(pInput) {
			// p2 found
			p2Found = true
			break

		} else {
			fmt.Sscanf(pInput[currentPos], "%s", &foundInstruction)

			if foundInstruction == "acc" {
				// get acc instruction
				fmt.Sscanf(pInput[currentPos], "acc%d", &accAmount)
				usedInstructions[currentPos] = true
				p1accumulator = p1accumulator + accAmount
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
	}
	return p1accumulator, p2Found
}

func changeInstructions(p2Input []string) int {
	var returnfromP1 int
	//p2currentPos := 0
	var p2Result bool
	changedInstructions := make([]bool, len(p2Input))
	p2foundInstruction := ""
	p2nopAmount := ""
	p2jmpAmount := ""

	for i, j := range p2Input {
		if p2Result == false {
			if changedInstructions[i] == false {
				// find nop and jmp and change them
				fmt.Sscanf(j, "%s", &p2foundInstruction)
				if p2foundInstruction == "nop" {
					// get amount
					fmt.Sscanf(p2Input[i], "nop%s", &p2nopAmount)
					p2Input[i] = "jmp " + p2nopAmount
					changedInstructions[i] = true
					returnfromP1, p2Result = processOperations(p2Input)

				} else if p2foundInstruction == "jmp" {

					// get amount
					fmt.Sscanf(p2Input[i], "jmp%s", &p2jmpAmount)
					p2Input[i] = "nop " + p2jmpAmount
					changedInstructions[i] = true
					returnfromP1, p2Result = processOperations(p2Input)
				}

			} else {
				// instruction already changed, just process
				returnfromP1, p2Result = processOperations(p2Input)
			}
		} else {
			// p2 found
			break
		}
		p2foundInstruction = ""
		p2nopAmount = ""
		p2jmpAmount = ""
	}

	return returnfromP1
}

func main() {

	// Grab user choices
	inputFile, _ := inputFlags()
	// pull in puzzle input
	puzzleInput, _ := readFile(inputFile)

	// Part 1
	p1Answer, _ := processOperations(puzzleInput)
	fmt.Println("Part 1: ", p1Answer)

	// Part 2
	p2Answer := changeInstructions(puzzleInput)
	fmt.Println("Part 2: ", p2Answer)

}
