package main

import (
	"fmt"
)

// function to handle the loop through each of the values
func loopThroughSlice(puzzleSlice []int) {
	// start loop from position 1
FirstLoop:
	for i := 0; i < len(puzzleSlice); i++ {
		for j := 0; j < len(puzzleSlice); j++ {
			answer := puzzleSlice[i] + puzzleSlice[j]
			if answer == 2020 {
				finalanswer := puzzleSlice[i] * puzzleSlice[j]
				foundpos1 := puzzleSlice[i]
				foundpos2 := puzzleSlice[j]
				fmt.Println("First number", foundpos1)
				fmt.Println("Second number", foundpos2)
				fmt.Println("Answer is: ", finalanswer)
				break FirstLoop
			}
		}
	}
}
