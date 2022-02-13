package main

import "fmt"

func findPairs(nums []int, k int) int {
    if ((k < 0) || (len(nums) == 0)) {
        return 0
    }
    var count int

    m := make(map[int]int, len(nums))

    for _, value := range nums {
        m[value]++
    }
    for key := range m {
        if ((k == 0) && (m[key] > 1)) {
            count++
            continue
        }
        if ((k > 0) && ((m[key + k]) > 0)) {
            count++
        }
    }

    return count
}

func main() {

    arr := [] int {3, 1, 4, 1, 5}
    arr1 := [] int {1, 2, 3, 4, 5}
    arr2 := [] int {1, 3, 1, 5, 4}
    arr3 := [] int {}

    fmt.Println(findPairs(arr, 2))
    fmt.Println(findPairs(arr1, 1))
    fmt.Println(findPairs(arr2, 0))
    fmt.Println(findPairs(arr3, 3))

}
