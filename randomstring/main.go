package main

import (
	"fmt"
	"math/rand"
)

func helloString() {
}

func randomString (l int) string {
    var letter = []rune("?/|&abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

    b := make([]rune, l)
    for i := range b {
        b[i] = letter[rand.Intn(len(letter))]
    }
    return string(b)
}

func main() {
    fmt.Printf("This testFunc to generate random strings\n")
    fmt.Printf("10: %s\n", randomString(10))
}
