package main

import (
	"bufio"
	"flag"
	"os"
	"sort"
)

// read a file from an input and return into a slice of strings
func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// use flags to get user input
func inputFlags() (string, string) {

	// variables declaration
	var flagfile string
	var partid string

	// flags declaration using flag package
	flag.StringVar(&flagfile, "file", "input.txt", "filename of the input data eg: input.txt")
	flag.StringVar(&partid, "part", "a", "use either a or b")

	flag.Parse() // after declaring flags we need to call it
	return flagfile, partid
}

// function to find seat IDs
func findSeatIDs(seats []string) (int, []int) {

	// Instructions to find seat:
	// For example, consider just the first seven characters of FBFBBFFRLR://
	// Start by considering the whole range, rows 0 through 127.
	// F means to take the lower half, keeping rows 0 through 63.
	// B means to take the upper half, keeping rows 32 through 63.
	// F means to take the lower half, keeping rows 32 through 47.
	// B means to take the upper half, keeping rows 40 through 47.
	// B keeps rows 44 through 47.
	// F keeps rows 44 through 45.
	// The final F keeps the lower of the two, row 44.

	// pseudo code
	// for loop each line and process against instructions above into two variables, row and column
	// multiply row by 8 and then add the column to get Seat ID
	// Store seat ID in a variable and on each run check if the new seat ID is higher than the last and overwrite if it is
	// Print out final highest seat ID

	maxSeatID := 0
	rowID := 0
	columnID := 0
	checkmaxSeatID := 0

	// for part b, create slice to store all seat IDs
	var seatIDslice []int

	for _, line := range seats {
		minRow := 0
		maxRow := 127
		minColumn := 0
		maxColumn := 7
		// F = 70
		// B = 66
		// R = 82
		// L = 76

		// process rows
		if line[0] == 70 {
			rowID = (minRow + maxRow) / 2
			maxRow = rowID
		} else {
			rowID = (minRow+maxRow)/2 + 1
			minRow = rowID
		}
		if line[1] == 70 {
			rowID = (minRow + maxRow) / 2
			maxRow = rowID
		} else {
			rowID = (minRow+maxRow)/2 + 1
			minRow = rowID
		}
		if line[2] == 70 {
			rowID = (minRow + maxRow) / 2
			maxRow = rowID
		} else {
			rowID = (minRow+maxRow)/2 + 1
			minRow = rowID
		}
		if line[3] == 70 {
			rowID = (minRow + maxRow) / 2
			maxRow = rowID
		} else {
			rowID = (minRow+maxRow)/2 + 1
			minRow = rowID
		}
		if line[4] == 70 {
			rowID = (minRow + maxRow) / 2
			maxRow = rowID
		} else {
			rowID = (minRow+maxRow)/2 + 1
			minRow = rowID
		}
		if line[5] == 70 {
			rowID = (minRow + maxRow) / 2
			maxRow = rowID
		} else {
			rowID = (minRow+maxRow)/2 + 1
			minRow = rowID
		}
		if line[6] == 70 {
			rowID = (minRow + maxRow) / 2
			maxRow = rowID
		} else {
			rowID = (minRow+maxRow)/2 + 1
			minRow = rowID
		}
		//fmt.Println(rowID)

		// process columns
		if line[7] == 76 {
			columnID = (minColumn + maxColumn) / 2
			maxColumn = columnID
		} else {
			columnID = (minColumn+maxColumn)/2 + 1
			minColumn = columnID
		}
		if line[8] == 76 {
			columnID = (minColumn + maxColumn) / 2
			maxColumn = columnID
		} else {
			columnID = (minColumn+maxColumn)/2 + 1
			minColumn = columnID
		}
		if line[9] == 76 {
			columnID = (minColumn + maxColumn) / 2
			maxColumn = columnID
		} else {
			columnID = (minColumn+maxColumn)/2 + 1
			minColumn = columnID
		}
		//fmt.Println(columnID)

		// calculate seat ID and check if its the highest
		checkmaxSeatID = ((rowID * 8) + columnID)

		// for part a
		if checkmaxSeatID > maxSeatID {
			maxSeatID = checkmaxSeatID
		}

		// for part b
		// add all seat ids to a slice
		seatIDslice = append(seatIDslice, checkmaxSeatID)
	}

	return maxSeatID, seatIDslice
}

func findmySeat(seatsToSort []int) int {
	//sort the slice into ascending order
	sort.Ints(seatsToSort)

	// loop through and find missing seat number
	for i := 1; i < len(seatsToSort)-1; i++ {
		if (seatsToSort[i+1] - seatsToSort[i]) == 2 {
			// seat ID found
			foundmySeat := (seatsToSort[i+1] + seatsToSort[i]) / 2
			return foundmySeat

		}
	}
	return 0
}
