package main

import (
    "fmt"
)

func maximizingXor(l int32, r int32) int32 {
    // Write your code here
    n := l ^ r
    result := int32(1)

    fmt.Printf("n : %v\n", n)
    for n > 0 {
        result <<= 1 //exponentiatlly grow/ power of two
        n >>= 1 //deivided by two
        fmt.Printf("n : %v\n", n)
        fmt.Printf("result : %v\n", result)
    }
    return result - 1
}

func main() {
    fmt.Println(maximizingXor(11, 12))
}
