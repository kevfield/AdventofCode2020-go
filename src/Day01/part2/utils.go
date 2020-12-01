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
			for k := 0; k < len(puzzleSlice); k++ {
				answer := puzzleSlice[i] + puzzleSlice[j] + puzzleSlice[k]
				if answer == 2020 {
					finalanswer := puzzleSlice[i] * puzzleSlice[j] * puzzleSlice[k]
					foundpos1 := puzzleSlice[i]
					foundpos2 := puzzleSlice[j]
					foundpos3 := puzzleSlice[k]
					fmt.Println("First number", foundpos1)
					fmt.Println("Second number", foundpos2)
					fmt.Println("Third number", foundpos3)
					fmt.Println("Answer is: ", finalanswer)
					break FirstLoop
				}
			}
		}
	}
}
