package main

import "fmt"

func findShortestSubArray(nums []int) int {
    frequency, maxFreq, subArray := map[int][]int{}, 0, len(nums)
    for i, num := range nums {
        if _, found := frequency[num]; !found {
            //[0] //num of occurs
            //[1] //storing start of the index
            //[2] //last index of occurs
            frequency[num] = [] int {1, i , i}
        } else {
            frequency[num][0]++
            frequency[num][2] = i
        }
        if maxFreq < frequency[num][0] {
            maxFreq = frequency[num][0]
        }
    }
    for _, indices := range frequency {
        if indices[0] == maxFreq {
            if (subArray > (indices[2] - indices[1] + 1)) {
                subArray = indices[2] - indices[1] + 1
            }
        }
    }
    return subArray

}

func main() {
    arr := [] int {1,2,2,3,1}
    arr2 := [] int {1,2,2,3,1,4,2}
    fmt.Println(findShortestSubArray(arr))
    fmt.Println(findShortestSubArray(arr2))
}
