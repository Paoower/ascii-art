package main

import (
	src "ascii-art/src"
	"fmt"
	"os"
	"strings"
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

	// Split the input by newline character
	segments := strings.Split(input, "\\n")
	finalOutput := make([]string, 0)

	for _, segment := range segments {
		lines := make([]string, 8)
		for _, char := range segment {
			c := strings.Split(src.GetLetter(content, int(char)), "\n")
			for i, l := range c {
				lines[i] += l
			}
		}
		finalOutput = append(finalOutput, lines...)
	}

	for _, line := range finalOutput {
		fmt.Println(line)
	}
}
