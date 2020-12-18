package main

import (
	"bufio"
	"flag"
	"os"
	"strconv"
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
	//lines = append(lines, "")
	return lines, scanner.Err()
}

// use flags to get user input
func inputFlags() (string, string, int) {

	// variables declaration
	var flagfile string
	var partid string
	var timestoRun int

	// flags declaration using flag package
	flag.StringVar(&flagfile, "file", "input.txt", "filename of the input data eg: input.txt")
	flag.StringVar(&partid, "part", "a", "use either a or b")
	flag.IntVar(&timestoRun, "timestorun", 1, "specify a number of times to run")

	flag.Parse() // after declaring flags we need to call it
	return flagfile, partid, timestoRun
}

func convertInputtoInt(strInput []string) []int {
	var intSlice []int
	var tmpstringSlice []string

	tmpstringSlice = strings.Split(strInput[0], ",")

	for a := 0; a < len(tmpstringSlice); a++ {
		b, _ := strconv.Atoi(tmpstringSlice[a])
		intSlice = append(intSlice, b)
	}

	//// sort the slice
	//sort.Ints(intSlice)

	return intSlice
}
