package main

import (
    "fmt"
)

func balancedSums(arr []int32) string {
    // Write your code here
    j := len(arr) - 1 //last index of array
    i := 0 //start index
    out := "NO"

    if (len(arr) <= 1) {
        return "YES"
    }

    left_sum := arr[i]
    right_sum := arr[j]
    for ; ((j - i) != 1); {
        //in the case of [1, 4, 1]
        if (((j - i) == 2) && (left_sum == right_sum)) {
            fmt.Printf(">>> \n")
            out = "YES"
            break
        }
        fmt.Printf("left-sum: %v right-sum: %v i: %v j: %v\n",
            left_sum, right_sum, i, j)
        if (left_sum > right_sum) {
            j--
            right_sum += arr[j]
        } else {
            i++
            left_sum += arr[i]
        }
    }
    if (left_sum == 0 || right_sum == 0) {
        out = "YES"
    }
    //fmt.Printf(">> \n", out)
    return out
}


func main() {
    //var arr = [] int32 {5,6,8,11}
    //var arr2 = [] int32 {1,2,3}
    var arr3 = [] int32 {1,2,3,3}

    //fmt.Println(balancedSums(arr))
    //fmt.Println(balancedSums(arr2))
    fmt.Println(balancedSums(arr3))

}
