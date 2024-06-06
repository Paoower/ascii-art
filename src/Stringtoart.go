package src

import (
	"strings"
)

// The GenerateASCIIArt function in Stringtoart.go is responsible for converting a given string input into ASCII art using a provided banner.
func Stringtoart(input string, content string) []string {
	segments := strings.Split(input, "\\n")
	finalOutput := make([]string, 0)

	for _, segment := range segments {
		lines := make([]string, 8)
		for _, char := range segment {
			c := strings.Split(GetLetter(content, int(char)), "\n")
			for i, l := range c {
				lines[i] += l
			}
		}
		finalOutput = append(finalOutput, lines...)
	}

	return finalOutput
}
