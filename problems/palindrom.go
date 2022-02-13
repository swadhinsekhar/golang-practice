package main

import "fmt"

func isPalindrome(inp string) bool {

    for i, j := 0, len(inp) - 1; i < j; i , j = i+1, j-1 {
        if (inp[i] == inp[j]) {
            continue
        } else {
            return false
        }
    }
    return true
}

func main() {
    fmt.Println(isPalindrome("MALAYALAM"))
    fmt.Println(isPalindrome("HELLO"))
}
