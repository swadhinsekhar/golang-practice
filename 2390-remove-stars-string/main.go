package main

import "fmt"

/*
2390. Removing Stars From a String
You are given a string s, which contains stars *.

In one operation, you can:

Choose a star in s.
Remove the closest non-star character to its left, as well as remove the star itself.
Return the string after all stars have been removed.

Note:

The input will be generated such that the operation is always possible.
It can be shown that the resulting string will always be unique.


Example 1:

Input: s = "leet**cod*e"
Output: "lecoe"
Explanation: Performing the removals from left to right:
- The closest character to the 1st star is 't' in "leet**cod*e". s becomes "lee*cod*e".
- The closest character to the 2nd star is 'e' in "lee*cod*e". s becomes "lecod*e".
- The closest character to the 3rd star is 'd' in "lecod*e". s becomes "lecoe".
There are no more stars, so we return "lecoe".


Example 2:

Input: s = "erase*****"
Output: ""
Explanation: The entire string is removed, so we return an empty string.
*/

// what is the time complexity?
// / The time complexity of the provided code is O(n), where n is the length of the input string s.
// / This is because the code iterates through each character in the string once,
// / performing constant time operations for each character. Therefore, the time complexity is linear with respect to the input size.
//
// / The space complexity of the provided code is O(n), where n is the length of the input string s.
// / This is because the code uses a stack to store non-star characters, which can grow to contain all characters in the string.
// / Additionally, the code uses a result array to store the final string, which can also grow to contain all characters in the string.
// / Therefore, the space complexity is linear with respect to the input size.
func removeStars(s string) string {
	result := make([]rune, 0, len(s))
	stack := make([]rune, 0, len(s))

	for _, char := range s {
		if char != '*' {
			stack = append(stack, char) // Push non-star characters onto the stack
			fmt.Println("stack", string(stack))
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1] // Remove the character closest to the left of the star
			fmt.Println("stack", string(stack))
		}
	}

	result = append(result, stack...) // Build the result string from the stack
	return string(result)
}

func main() {
	// Code
	s := "leet**cod*e"
	fmt.Println("lecoe", removeStars(s))
	//s = "erase*****"
	//fmt.Println("", removeStars(s))
}
