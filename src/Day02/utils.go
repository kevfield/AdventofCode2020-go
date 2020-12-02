package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
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

// check password function to confirm which passwords are valid
func checkPasswordsA(passwordlist []string) int {

	// split the input into sections
	// - min/max values
	// - password character value
	// - password string
	//
	// loop through each line and check for passwords that meet the criteria
	// add to a correct count for each correcrt password
	// return correct count

	var minvalue int
	var maxvalue int
	var passchar string
	var fullpassword string

	// initialise correct password count and set to 0
	correctpasswordsA := 0

	for _, inputfileline := range passwordlist {
		fmt.Sscanf(inputfileline, "%d-%d %s %s", &minvalue, &maxvalue, &passchar, &fullpassword)
		// remove : character
		passchartrimmed := strings.Replace(passchar, ":", "", -1)
		// start counting
		countoccurances := strings.Count(fullpassword, passchartrimmed)
		if countoccurances >= minvalue && countoccurances <= maxvalue {
			correctpasswordsA = correctpasswordsA + 1
		}
	}

	return correctpasswordsA
}

func checkPasswordsB(passwordlist []string) int {

	var minvalue int
	var maxvalue int
	var passchar rune
	var fullpassword string

	// initialise correct password count and set to 0
	correctpasswordsB := 0

	for _, inputfileline := range passwordlist {
		fmt.Sscanf(inputfileline, "%d-%d %c: %s", &minvalue, &maxvalue, &passchar, &fullpassword)

		if fullpassword[minvalue-1] == byte(passchar) && fullpassword[maxvalue-1] != byte(passchar) {
			correctpasswordsB = correctpasswordsB + 1
		} else if fullpassword[minvalue-1] != byte(passchar) && fullpassword[maxvalue-1] == byte(passchar) {
			correctpasswordsB = correctpasswordsB + 1
		}

	}

	return correctpasswordsB
}
