package main

import (
    "fmt"
    "sort"
)

func missingNumbers(arr []int32, brr []int32) []int32 {
    // Write your code here
    var res = [] int32{}

    var c1 = make(map[int32]int32)
    var c2 = make(map[int32]int32)

    for _, v := range arr {
        if _, ok := c1[int32(v)]; ok {
            c1[int32(v)]++
        } else {
            c1[int32(v)] = 1
        }
    }
    for _, v := range brr {
        if _, ok := c2[int32(v)]; ok {
            c2[int32(v)]++
        } else {
            c2[int32(v)] = 1
        }
    }

    for k, v := range c2 {
        if val, ok := c1[k]; ok {
            if ((val - v) != 0) {
                res = append(res, k)
            }
        } else {
            res = append(res, k)
        }
    }

    sort.Slice(res, func(i, j int) bool {
        return res[i] < res[j]
    })
    return res
}

func main() {
    var arr = [] int32 {203, 204, 205, 206, 207, 208, 203, 204, 205, 206}
    var brr = [] int32 {203, 204, 204, 205, 206, 207, 205, 208, 203, 206, 205, 206, 204}

    fmt.Println(missingNumbers(arr, brr))

}
