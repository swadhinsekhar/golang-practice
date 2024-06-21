package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	//check length
	if len(word1) < 1 {
		return word1 + word2
	}
	i := 0
	j := 0
	res := []rune{}
	for i < len(word1) && j < len(word2) {
		res = append(res, rune(word1[i]))
		res = append(res, rune(word2[i]))
		i = i + 1
		j = j + 1
	}
	//left letters
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

func rotatedArray(arr []int) []int {
	l := len(arr) - 1
	for i, j := 0, l; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdOfStrings(str1, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[:gcd(len(str1), len(str2))]
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println(rotatedArray(arr))
	word1 := "abc"
	word2 := "pqr"
	res := "apbqcr"
	fmt.Println(mergeAlternately(word1, word2) == res)
	str1 := "ABCABC"
	str2 := "ABC"
	res = "ABC"
	fmt.Println(gcdOfStrings(str1, str2) == res)
	str1 = "ABABAB"
	str2 = "ABAB"
	res = "AB"
	fmt.Println(gcdOfStrings(str1, str2) == res)
	str1 = "LEET"
	str2 = "CODE"
	fmt.Println(gcdOfStrings(str1, str2) == "")

}
