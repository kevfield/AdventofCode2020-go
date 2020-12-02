package main

import "fmt"

func main() {

	// get the file and part input from the user
	inputfile, userpart := inputFlags()

	// get password list using file specified by user
	passwordlist, _ := readFile(inputfile)

	if userpart == "a" {

		puzzleanswer := checkPasswords(passwordlist)
		fmt.Println("Amount of correct passwords in list: ", puzzleanswer)

	} else {

		fmt.Println("not coded yet")

	}

}
