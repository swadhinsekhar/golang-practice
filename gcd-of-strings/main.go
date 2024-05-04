package main

import "fmt"

//[1071. Greatest Common Divisor of Strings](https://leetcode.com/problems/greatest-common-divisor-of-strings/)

// Time Complexity: O(log(min(n, m)))
// Space Complexity: O(1)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[:gcd(len(str1), len(str2))]
}

func main() {
	// Code
	fmt.Println("Hello, World!")

	str1 := "ABCABC"
	str2 := "ABC"
	// Output:** "ABC"
	fmt.Println("ABC", gcdOfStrings(str1, str2))

	str1 = "ABABAB"
	str2 = "ABAB"
	// Output:** "AB"
	fmt.Println("AB", gcdOfStrings(str1, str2))

	str1 = "LEET"
	str2 = "CODE"
	// Output: ""
	fmt.Println("", gcdOfStrings(str1, str2))
}
