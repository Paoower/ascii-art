package main

import (
	"fmt"
	src "ascii-art/src"
)

func main() {
	content := src.FileOpen("banners/shadow.txt")
	fmt.Println(src.GetLetter(content, 97))
}