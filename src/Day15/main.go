package main

import "fmt"

func elfGame(gameInput []int, gameAmount int) {
	//make map to track numbers
	// key = number from elf
	// value = turn number
	rememberStuff := make(map[int]int)
	//currNumber := 0
	//nextNumber := 0
	currNumber := 0
	//var prevTurn int

	//// setup map with already answered numbers and the following number without processing in order to start the full loop
	//for item := 0; item < len(gameInput); item++ {
	//	rememberStuff[gameInput[item]] = item
	//}
	//if rememberStuff

	for turn := 0; turn <= gameAmount; turn++ {
		// starting number checks
		if turn < len(gameInput) {
			// starting number found, add it to map
			rememberStuff[gameInput[turn]] = turn
			currNumber = gameInput[turn]
			//fmt.Println("turn", turn)
			//fmt.Println("curr", gameInput[turn])
		} else {
			// calc next number
			_, ok := rememberStuff[currNumber]
			if ok == false {
				// new number found so next number is 0
				// update map and reset preNumber
				rememberStuff[currNumber] = turn
				currNumber = 0

			} else {
				// if number already spoken then new number is how many turns apart it was spoken previously
				currNumber = turn - rememberStuff[currNumber]
				fmt.Println("turn", turn)
				fmt.Println("prev turn", rememberStuff[currNumber])
				fmt.Println(currNumber)
				rememberStuff[currNumber] = turn
			}
		}

	}
	//fmt.Println("current", currNumber)
	fmt.Println("++++++++++++++++++++++++++++++++")
	fmt.Println(rememberStuff)
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
