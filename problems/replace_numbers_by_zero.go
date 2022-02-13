package main

import "fmt"

func main() {
    str := "abc123465@ahsjf222"
    newStr := make([] rune, len(str))

    i, added := 0, false

    for _, r := range str {
        if r >= '0' && r <= '9' {
            if added {
                continue
            }
            added, newStr[i] = true, '0'
        } else {
            added, newStr[i] = false, r
        }
        i++
    }
    fmt.Println(string(newStr[:i]))
}
