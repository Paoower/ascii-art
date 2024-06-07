package main_test

import (
	src "ascii-art/src"
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	correct := []string{" _    _          _   _          ", "| |  | |        | | | |         ", "| |__| |   ___  | | | |   ___   ", "|  __  |  / _ \\ | | | |  / _ \\  ", "| |  | | |  __/ | | | | | (_) | ", "|_|  |_|  \\___| |_| |_|  \\___/  ", "                                ", "                                "}
	input := "Hello"
	lines := make([]string, 0)
	words := strings.Split(input, "\\n")

	for _, word := range words {
		lines = append(lines, src.GetWord(word)...)
	}

	for i, line := range lines {
		if correct[i] != line {
			t.Errorf("Program Output: %v \n", lines)
			t.Errorf("Correct Output: %v \n", correct)
		}
	}
}

func TestHelloThere(t *testing.T) {
	correct := []string{" _    _          _   _               _______   _                           ", "| |  | |        | | | |             |__   __| | |                          ", "| |__| |   ___  | | | |   ___          | |    | |__     ___   _ __    ___  ", "|  __  |  / _ \\ | | | |  / _ \\         | |    |  _ \\   / _ \\ | '__|  / _ \\ ", "| |  | | |  __/ | | | | | (_) |        | |    | | | | |  __/ | |    |  __/ ", "|_|  |_|  \\___| |_| |_|  \\___/         |_|    |_| |_|  \\___| |_|     \\___| ", "                                                                           ", "                                                                           "}
	input := "Hello There"
	lines := make([]string, 0)
	words := strings.Split(input, "\\n")

	for _, word := range words {
		lines = append(lines, src.GetWord(word)...)
	}

	for i, line := range lines {
		if correct[i] != line {
			t.Errorf("Program Output: %v \n", lines)
			t.Errorf("Correct Output: %v \n", correct)
		}
	}
}
