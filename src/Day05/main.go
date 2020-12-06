package main

import (
	"fmt"
)

func main() {

	fileLocation, flag := inputFlags()

	puzzleInput, _ := readFile(fileLocation)

	if flag == "a" {

		highestSeatID, _ := findSeatIDs(puzzleInput)
		fmt.Println(highestSeatID)

	} else if flag == "b" {

		_, seatSlice := findSeatIDs(puzzleInput)
		mySeat := findmySeat(seatSlice)
		fmt.Println(mySeat)

	} else {

		fmt.Println("Invalid part choice, please specify either a or b")

	}

}
