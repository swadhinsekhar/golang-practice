package main

import (
	"fmt"
)

func mergeAlternately(word1 string, word2 string) string {
	if len(word1) < 1 {
		return word1 + word2
	}
	//if len(word2) > 100 {
	//    return ""
	//}
	i, j := 0, 0
	res := []rune{}
	for i < len(word1) && j < len(word2) {
		res = append(res, rune(word1[i]))
		res = append(res, rune(word2[j]))
		i += 1
		j += 1
	}
	for i < len(word1) {
		res = append(res, rune(word1[i]))
		i++
	}
	for j < len(word2) {
		res = append(res, rune(word2[j]))
		j++
	}

	return string(res)
}

func main() {
	word1 := "abc"
	word2 := "pqr"
	fmt.Println(mergeAlternately(word1, word2))
}
