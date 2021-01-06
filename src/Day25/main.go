package main

import "fmt"

func getpublicKeys(keysInput []int) (int, int) {
	cardPub := keysInput[0]
	doorPub := keysInput[1]

	return cardPub, doorPub
}

func findLoopSize(publicKey int) int {

	loopSize := 0
	loopAmount := 1

	for loopAmount != publicKey {
		loopSize++
		loopAmount = (loopAmount * 7) % 20201227
		//fmt.Println(loopSize)
	}

	return loopSize

}

func calcencryptionKeys(cardLoop int, doorLoop int, cardPK int, doorPK int) int {

	encKey := 1

	for i := 1; i <= cardLoop; i++ {
		encKey = (encKey * doorPK) % 20201227
	}

	return encKey
}

func main() {

	// get input from user
	fileLoc, _ := inputFlags()

	// get puzzle input
	puzzleinput, _ := readFile(fileLoc)

	// convert puzzle input to slice of Int
	intInput := convertInputtoInt(puzzleinput)

	// get the public keys
	cardpublicKey, doorpublicKey := getpublicKeys(intInput)
	//fmt.Println(cardpublicKey, doorpublicKey)

	// find the loop sizes for each key
	cardloopSize := findLoopSize(cardpublicKey)
	doorloopSize := findLoopSize(doorpublicKey)

	// get the encryption keys
	encryptionKey := calcencryptionKeys(cardloopSize, doorloopSize, cardpublicKey, doorpublicKey)
	fmt.Println("Part 1: ", encryptionKey)
}
