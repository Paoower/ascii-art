package src

import (
	"strings"
)

func GetWord(input string) []string {
	lines := make([]string, 8)
	content := FileOpen("banners/standard.txt")

	for _, char := range input {
		c := strings.Split(GetLetter(content, int(char)), "\n")
		for i, l := range c {
			lines[i] += l
		}
	}
	return lines
}
