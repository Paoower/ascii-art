package src

import (
	"strings"
)

const LETTER_HEIGHT = 8

func GetLetter(content string, ascii int) string {
	str := ""
	lines := strings.Split(content, "\n")

	place := ascii - 31
	times := (place - 1) * LETTER_HEIGHT
	beginning := (ascii - 30) + times

	for i := beginning; i < beginning+LETTER_HEIGHT; i++ {
		if i != (beginning+LETTER_HEIGHT)-1 {
			str += lines[i-1] + "\n"
		} else {
			str += lines[i-1]
		}
	}

	return str
}
