package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	banner := make(map[rune][]string)

	for i := 0; i <= 94; i++ {
		char := rune(32 + i)
		start := i*9 + 1
		if start+8 > len(lines) {
			return nil, errors.New("banner file format incorrect")
		}
		banner[char] = lines[start : start+8]
	}
	return banner, nil
}
