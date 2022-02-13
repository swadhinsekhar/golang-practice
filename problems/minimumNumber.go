package main

import  (
    "fmt"
_    "regexp"
)

func isCapLetter(str string) bool {
    for _, c := range str {
        if (c >= 'A' && c <= 'Z') {
            return true
        }
    }
    return false
}

func isSmallLetter(str string) bool {
    for _, c := range str {
        if ('a' <= c && c <= 'z') {
            return true
        }
    }
    return false
}

func isNumber(str string) bool {
    for _, c := range str {
        if ('0' <= c && c <= '9') {
            return true
        }
    }
    return false
}

func isSpecialChar(str string) bool {
    special_characters := "!@#$%^&*()-+"

    for _, c := range str {
        for _, s := range special_characters {
            if c == s {
                return true
            }
        }
    }
    return false
}

func minimumNumber(n int32, password string) int32 {
    // Return the minimum number of characters to make the password strong
    //numbers := "0123456789"
    //lower_case := "abcdefghijklmnopqrstuvwxyz"
    //upper_case := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    //special_characters := "!@#$%^&*()-+"
    min_len := 6
    pwd_len := int(n)
    count := 0

    fmt.Println(n, password)
    if (pwd_len < min_len) {
        return (int32(min_len) - n)
    }

    if ret := isCapLetter(password); ret != true {
        count++
    }
    if ret := isSmallLetter(password); ret != true {
        count++
    }
    if ret := isNumber(password); ret != true{
        count++
    }

    if ret := isSpecialChar(password); ret != true{
        count++
    }

    rest := int32(min_len - len(password))
    if ((rest > 0) && (rest >= count)) {
        return rest
    } else {
        return count
    }

}

func main() {
    pwd := "#HackerRank"

    fmt.Println(minimumNumber(int32(len(pwd)), pwd))

}
