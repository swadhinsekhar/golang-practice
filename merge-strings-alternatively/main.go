package main

import "fmt"

//[1768. Merge Strings Alternately](https://leetcode.com/problems/merge-strings-alternately/)

func MergeAlneternately(word1 string, word2 string) string {
	if len(word1) < 1 {
		return word1 + word2
	}
	// Code
	var result string
	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		result += string(word1[i])
		result += string(word2[i])
		i++
		j++
	}

	for i < len(word1) {
		result += string(word1[i])
		i++
	}
	for j < len(word2) {
		result += string(word2[j])
		j++
	}

	return result
}

func main() {
	// Code
	fmt.Println("Hello, World!")
	word1 := "abc"
	word2 := "pqr"
	// Output: "apbqcr"
	fmt.Println(MergeAlneternately(word1, word2))

	word1 = "ab"
	word2 = "pqrs"
	fmt.Println(MergeAlneternately(word1, word2))
	// Output: "apbqrs"
}
