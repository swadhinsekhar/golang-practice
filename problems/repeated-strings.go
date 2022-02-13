package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

func repeatedString(s string, n int64) int64 {
    // Write your code here
    inputLength := int64(len(s))
	occurrenceOfAInInput := occurrenceOfA(s)
	numberOfSubstring := int64(n / inputLength)

	numberOfAs := numberOfSubstring  * occurrenceOfAInInput
	reminder := n % inputLength // a

	for i := int64(0); i < reminder; i++ {
		if s[i]  == 'a' {
			numberOfAs++
		}
	}

	return numberOfAs
}

func occurrenceOfA(s string) int64  {
	var count int64
	for i := range s {
		if s[i] == 'a' {
			count++
		}
	}

	return count
}

func main() {

    s := 'abaabaaba'

    result := repeatedString(s, n)
}

