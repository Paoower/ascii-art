package main

import (
	src "ascii-art/src"
	"fmt"
	"os"
)

func main() {

	// Checking if the correct number of arguments is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <string>")
		return
	}
	// Get the input string from command line arguments
	input := os.Args[1]
	// Load the standard banner file
	bannerFile := "banners/standard.txt"
	content := src.FileOpen(bannerFile)
	// Generate ASCII art
	finalOutput := src.Stringtoart(input, content)
	// Print the final output
	for _, line := range finalOutput {
		fmt.Println(line)
	}
}
