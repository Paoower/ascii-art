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
		fmt.Println("Wrong use of the tool")
		return
	}
	// Get the input string from command line arguments
	input := os.Args[1]

	lines := make([]string, 0)
	words := strings.Split(input, "\\n")

	for _, word := range words {
		lines = append(lines, src.GetWord(word)...)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
