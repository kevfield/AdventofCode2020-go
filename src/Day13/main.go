package main

import (
	"fmt"
	"strconv"
	"strings"
)

func busTimes(p1Input []string) int {

	var calculatedTime int
	var returnedbusNo int
	var foundBus int
	earliestbusTime := 0
	// pull in time stamp from input file abnd convert to string
	timeStamp := p1Input[0]
	convtimeStamp, _ := strconv.Atoi(p1Input[0])
	// make new slice for the buses
	var busSlice []string
	busSlice = strings.Split(p1Input[1], ",")

	for _, busNo := range busSlice {
		if busNo == timeStamp {
			// skip
		} else if busNo != "x" {
			bustoSend, _ := strconv.Atoi(busNo)
			calculatedTime, returnedbusNo = busAlg(convtimeStamp, bustoSend)
		}
		if earliestbusTime == 0 {
			// first run
			earliestbusTime = earliestbusTime + calculatedTime
		} else {
			if calculatedTime < earliestbusTime {
				earliestbusTime = calculatedTime
				foundBus = returnedbusNo
			}
		}
	}
	busCalc := earliestbusTime - convtimeStamp
	p1result := busCalc * foundBus

	return p1result
}

func busAlg(timetoCalc int, busNo int) (int, int) {
	// run the bus number calcs
	step1 := timetoCalc / busNo
	step2 := busNo * step1
	for step2 < timetoCalc {
		step2 = step2 + busNo
	}
	return step2, busNo
}

func main() {

	// get input from user
	fileLoc, _ := inputFlags()

	// get puzzle input
	puzzleinput, _ := readFile(fileLoc)

	fmt.Println("Part 1: ", busTimes(puzzleinput))

}
