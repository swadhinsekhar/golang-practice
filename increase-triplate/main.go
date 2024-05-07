package main

import "fmt"

/*
Input: nums = [1,2,3,4,5]
Output: true
Explanation: Any triplet where i < j < k is valid.
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
