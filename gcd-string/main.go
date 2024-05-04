package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {
	// Find the lengths of the strings
	len1, len2 := len(str1), len(str2)

	// Find the greatest common divisor (GCD) of the lengths
	gcd := gcd(len1, len2)

	// Check if the repeated substring forms the other string
	if str1[:gcd] == str2[:gcd] && str1 == repeat(str1[:gcd], len1/gcd) && str2 == repeat(str1[:gcd], len2/gcd) {
		return str1[:gcd]
	}

	return ""
}

// Function to find the greatest common divisor (GCD) of two numbers
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to repeat a string 's' 'n' times
func repeat(s string, n int) string {
	res := ""
	for i := 0; i < n; i++ {
		res += s
	}
	return res
}

func main() {
	str1 := "ABCABC"
	str2 := "ABC"
	fmt.Println(gcdOfStrings(str1, str2)) // Output: "ABC"
}
