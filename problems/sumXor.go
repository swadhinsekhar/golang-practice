package main

import (
	"fmt"
)

func findMSB(n int64) int64 {

    var i int64
    for i = 64; i >= 0; i-- {
        if val := (n & (1 << i)); val != 0 {
            return i
        }
    }
    return 0
}

func sumXor(n int64) int64 {
    // Write your code here
    var i, count int64
    var msb int64

    //find out msb of the number
    msb = findMSB(n)

    for i = 0; i < msb; i++ {
        if ((n & (1 << i)) == 0) {
            count++
        }
    }
    return (1 << count)
}


func main() {
    fmt.Println(sumXor(5))
}
