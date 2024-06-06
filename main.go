package main

import (
	src "ascii-art/src"
	"fmt"
)

func main() {
	content := src.FileOpen("banners/shadow.txt")
	test := "hello"

	for _, char := range test {
		fmt.Println(src.GetLetter(content, int(char)))
	}
}
