package main

import (
	"golang.org/x/tour/wc"
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {

	m := make(map[string]int)
	words := strings.Fields(s)
	
	for i := 0; i < len(words); i++ {
		fmt.Println(words[i])
		m[words[i]] += 1	
	}

	return m
}

func main() {

	wc.Test(WordCount)
}
