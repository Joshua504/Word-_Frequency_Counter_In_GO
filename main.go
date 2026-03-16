package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type kv struct{
	key string
	value int
}

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

	var sliceOfStruct []kv

	for i, k := range wordCount{
		sliceOfStruct = append(sliceOfStruct, kv{i, k})
	}

	sort.Slice(sliceOfStruct, func(i, j int) bool {
    	return sliceOfStruct[i].value > sliceOfStruct[j].value
	})

	var finalReslt string

	for _, i := range sliceOfStruct {
		line := fmt.Sprintf("%s: %d\n", i.key, i.value)
		finalReslt += line
	}

	os.WriteFile("result.txt", []byte(finalReslt), 0644)
}