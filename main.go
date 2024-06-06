package main

import (
	src "ascii-art/src"
	"fmt"
	"os"
)

func main() {

	// Checking if the correct number of arguments is provided
	if len(os.Args) < 2 {
		fmt.Println("Wrong use of the tool")
		return
	}
	// Get the input string from command line arguments
	input := os.Args[1:]
	content := src.FileOpen("banners/shadow.txt")

	for _, word := range input {
		if word == "\n" {
			fmt.Println("\n")
			continue
		}
		for _, char := range word {
			fmt.Println(src.GetLetter(content, int(char)))
		}
		//Add a space after each word
		fmt.Println(src.GetLetter(content, 32))
	}
}
