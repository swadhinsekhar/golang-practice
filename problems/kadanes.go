package main

import "fmt"


/*
 * Given an integer array nums, find a contiguous non-empty
 * subarray within the array that has the largest product,
 * and return the product.
**/
func maxProduct(nums []int) int {
    curMax, curMin, ans := nums[0], nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        tmp := curMax
        curMax = max(nums[i], max(tmp * nums[i], curMin * nums[i]))
        curMin = min(nums[i], min(tmp * nums[i], curMin * nums[i]))
        if curMax > ans { ans = curMax }
    }
    return ans
}

/* kadan's algrorithm */
func maxSumSubArray(nums [] int) int {
    var maxSum, curSum int = 0, 0

    for i := 0; i < len(nums); i++ {

        curSum = curSum + nums[i]

        if (curSum > maxSum) {
            maxSum = curSum
        }

        //skipping -ve numbers
        if (curSum < 0) {
            curSum = 0
        }
    }
    return maxSum
}

func max(a, b int) int {
    if a > b { return a }
    return b
}

func min(a, b int) int {
    if a < b { return a }
    return b
}

func main() {

    arr := [] int {2,3,-2,4}
    arr2 := [] int {5, -4, -2, 6, 6, -1}

    fmt.Println(maxProduct(arr))
    fmt.Println(maxSumSubArray(arr2))
}
