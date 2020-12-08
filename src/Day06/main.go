package main

import "fmt"

func processLandingCards(input []string) (int, int) {

	var groupsCards map[byte]int
	var p1countResponses int
	var p1totalResponses int
	var p2countResponses int
	var p2totalResponses int
	newPerson := 0

	groupsCards = make(map[byte]int)

	for _, person := range input {
		if person == "" {
			// reset count responses
			p1countResponses = 0
			p2countResponses = 0
			for _, value := range groupsCards {
				if value >= 1 {
					p1countResponses++
				}
			}
			for _, value := range groupsCards {
				if value >= newPerson {
					p2countResponses++
				}
			}
			newPerson = 0
			p1totalResponses = p1totalResponses + p1countResponses
			p2totalResponses = p2totalResponses + p2countResponses
			// found new group
			// make new map
			groupsCards = make(map[byte]int)
		} else {
			newPerson++
		}
		// loop through lines to get answers and add to map
		for _, answer := range person {
			groupsCards[byte(answer)]++
		}

	}
	// edge case p1, count last line
	p1countResponses = 0
	for _, value := range groupsCards {
		if value >= 1 {
			p1countResponses++
		}
	}
	p1totalResponses = p1totalResponses + p1countResponses

	// edge p2
	p2countResponses = 0
	for _, value := range groupsCards {
		if value >= newPerson {
			p2countResponses++
		}
	}
	p2totalResponses = p2totalResponses + p2countResponses
	return p1totalResponses, p2totalResponses
}

func main() {

	// get input from file
	puzzleInput, _ := readFile("input.txt")

	// run part1
	p1answer, p2answer := processLandingCards(puzzleInput)
	fmt.Println("Part 1 =", p1answer)
	fmt.Println("Part 2 =", p2answer)

}
