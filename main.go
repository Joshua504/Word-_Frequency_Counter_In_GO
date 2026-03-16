package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("text.txt")

	if err != nil {
		fmt.Println("error: FILE DOES NOT EXIST")
		return
	}

	splitWord := strings.Fields(string(content))

	wordCount := make(map[string]int)

	for _, word := range splitWord{
		wordCount[word] ++

	}

	fmt.Println(wordCount)
}