package main

import (
	"strings"
)

func GenerateAsciiArt(input string, banner map[rune][]string) string {
	lines := strings.Split(input, "\n")
	result := []string{}
	for _, line := range lines {
		asciiLines := make([]string, 8)

		for _, char := range line {
			art, ok := banner[char]
			if !ok {
				art = banner[' ']
			}
			for i := 0; i < 8; i++ {
				asciiLines[i] += art[i]
			}
		}
		result = append(result, asciiLines...)
	}
	return strings.Join(result, "\n")
}
