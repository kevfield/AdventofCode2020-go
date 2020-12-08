package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
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

func processLandingCards(input []string) {
	var newstuff string
	count := 0
	totalfoundmatches := 0
	foundMatches := 0
	partbCount := 0
	totalpartb := 0
	var dict map[rune]int

	for _, line := range input {
		if len(line) == 0 {
			partbCount = 0
			dict = make(map[rune]int, 0)
			for len(newstuff) > 0 {

				r, size := utf8.DecodeRuneInString(strings.ToLower(newstuff))
				//fmt.Printf("%c %v\n", r, size)
				newstuff = newstuff[size:]
				if dict[r] > 1 {
					continue
				}
				//fmt.Println(dict[r])
				if dict[r] == 1 {
					count++
				}
				dict[r]++
			}
			//part a
			foundMatches = len(dict)

			totalfoundmatches = totalfoundmatches + foundMatches
			totalpartb = totalpartb + partbCount

		} else {
			for _, i := range strings.Split(line, " ") {
				newstuff = newstuff + strings.Replace(i, "\n", "", -1)
			}
		}

	}

	fmt.Println("Part a: ", totalfoundmatches)
	fmt.Println("Part b: ", totalpartb)
}
