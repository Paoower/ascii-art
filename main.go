package main

import (
	src "ascii-art/src"
	"fmt"
)

func main() {
	content := src.FileOpen("banners/shadow.txt")
	fmt.Println(src.GetLetter(content, 98))
}
