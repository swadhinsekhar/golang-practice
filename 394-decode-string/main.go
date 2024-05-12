package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
394. Decode String
Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there will not be input like 3a or 2[4].

The test cases are generated so that the length of the output will never exceed 10^5.

Example 1:

Input: s = "3[a]2[bc]"
Output: "aaabcbc"
Example 2:

Input: s = "3[a2[c]]"
Output: "accaccacc"
Example 3:

Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"

*/

func decodeString(s string) string {
	decoded := ""
	for _, char := range s {
		if char == ']' {
			decoded = decodeLast(decoded)
		} else {
			decoded += string(char)
		}
	}
	return decoded
}

func decodeLast(encoded string) string {
	start := strings.LastIndex(encoded, "[")
	sub := encoded[start+1:]
	end := start
	start--

	for start >= 0 && unicode.IsDigit(rune(encoded[start])) {
		start--
	}

	rep, _ := strconv.Atoi(encoded[start+1 : end])
	return encoded[:start+1] + strings.Repeat(sub, rep)
}

func main() {
	fmt.Println("decode-string")
	s := "3[a]2[bc]"
	fmt.Println("aaabcbc", decodeString(s))
}
