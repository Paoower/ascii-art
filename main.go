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
	content := src.FileOpen("banners/standard.txt")

	lines := make([]string, 8)

	for _, char := range input {
		c := strings.Split(src.GetLetter(content, int(char)), "\n")
		for i, l := range c {
			lines[i] += l
		}
	}
	for _, line := range lines {
		fmt.Println(line)
	}

}
