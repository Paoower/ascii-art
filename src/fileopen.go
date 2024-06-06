package src

import "os"

func FileOpen(filename string) string {
	f, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(f)
}
