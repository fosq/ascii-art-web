package main

import (
	"fmt"
	"os"
	"strings"
)

func getAscii(input string, font string) []string {
	f, err := os.ReadFile("ascii/" + font)
	check(err)

	// Split ascii datafile by newline
	data := strings.Split(string(f), "\n")

	ascii := printAscii(data, input)
	return ascii
}

// Prints asciicode characters
func printAscii(textArr []string, input string) []string {
	var asciiArray []string
	splitInput := strings.Split(input, `\n`) // Split input by newline into a string array
	for i := range splitInput {              // Loop over splitInput
		if splitInput[i] != "" { // If string is not empty, continue
			for j := 0; j < 8; j++ { // Loop over 8, the height of an ascii character
				var printableText string                // String variable for printing ascii characters line by line
				for _, element := range splitInput[i] { // Loop by element in splitInput[i]
					if element > 31 && element < 127 { // If element not ascii, skip
						printableText += textArr[((int(element)-32)*9)+j+1] // Add every row of needed ascii character into a string
					}
				}
				asciiArray = append(asciiArray, printableText)
			}
		}
	}
	return asciiArray
}

// Error checker
func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
