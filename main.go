package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . sample.txt | cat -e")
		return
	}
	input := os.Args[1]

	banner, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Println("error loading file", err)
		return
	}
	output := GenerateAsciiArt(input, banner)

	fmt.Println(output)
}
