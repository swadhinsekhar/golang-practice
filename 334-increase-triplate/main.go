package main

import "fmt"

/*
334. Increasing Triplet Subsequence
Given an integer array nums, return true if there exists a triple of
indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.

Input: nums = [1,2,3,4,5]
Output: true
Explanation: Any triplet where i < j < k is valid.

Example 2:

Input: nums = [5,4,3,2,1]
Output: false
Explanation: No triplet exists.
Example 3:

Input: nums = [2,1,5,0,4,6]
Output: true
Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.
*/

func increasingTriplet(nums []int) bool {
	first := int(^uint(0) >> 1) // Max unsigned int >> 1 is the max signed int, basically it divides the max unsigned int by 2
	second := int(^uint(0) >> 1)
	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(increasingTriplet(nums))

	nums = []int{5, 4, 3, 2, 1}
	fmt.Println(nums)
	fmt.Println(increasingTriplet(nums))
}
