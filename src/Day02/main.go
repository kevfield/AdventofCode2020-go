package main

import "fmt"

func main() {

	// get the file and part input from the user
	inputfile, userpart := inputFlags()

	// get password list using file specified by user
	passwordlist, _ := readFile(inputfile)

	if userpart == "a" {

		puzzleanswer := checkPasswordsA(passwordlist)
		fmt.Println("Amount of correct passwords in list: ", puzzleanswer)

	} else if userpart == "b" {

		output := checkPasswordsB(passwordlist)
		fmt.Println(output)

	} else {

		fmt.Println("Invalid part choice specified, options are a or b")

	}

}
