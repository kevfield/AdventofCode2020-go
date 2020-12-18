package main

import "fmt"

func elfGame(gameInput []int, gameAmount int) {
	//make map to track numbers
	// key = number from elf
	// value = turn number
	rememberStuff := make(map[int]int)
	currNumber := 0
	previousNumber := 0

	for turn := 0; turn < gameAmount; turn++ {
		// starting number checks
		if turn < len(gameInput) {
			// starting number found, add it to map
			if turn == len(gameInput)-1 {
				previousNumber = gameInput[turn]
			} else {
				rememberStuff[gameInput[turn]] = turn
			}
		} else {
			// calc next number
			_, ok := rememberStuff[previousNumber]
			//fmt.Println("Ok is:", ok)
			if ok == false {
				// new number found so next number is 0
				// update map and reset preNumber
				rememberStuff[previousNumber] = turn - 1
				previousNumber = 0
			} else {
				// if number already spoken then new number is how many turns apart it was spoken previously
				currNumber = turn - 1 - rememberStuff[previousNumber]
				rememberStuff[previousNumber] = turn - 1
				previousNumber = currNumber
			}
		}
	}
	fmt.Println(currNumber)

}
func main() {
	// get input from user
	fileLoc, _, runTimes := inputFlags()
	// get puzzle input
	puzzleinput, _ := readFile(fileLoc)
	// convert puzzle input to slice of Int
	intInput := convertInputtoInt(puzzleinput)
	elfGame(intInput, runTimes)
}
