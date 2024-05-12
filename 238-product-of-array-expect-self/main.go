package main

import (
	"fmt"
)

/*
238. Product of Array Except Self
Given an integer array nums, return an array answer such that answer[i] is equal
to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

**Example 1:**

**Input:** nums = [1,2,3,4]
**Output:** [24,12,8,6]

**Example 2:**

**Input:** nums = [-1,1,0,-3,3]
**Output:** [0,0,9,0,0]
*/

func productExceptSelf(nums []int) []int {
	n := len(nums)
	output := make([]int, n)
	left := make([]int, n)
	right := make([]int, n)

	/*
		The left array is used to store the product of all the numbers to the left of the current index in the nums array.
		It is initialized such that the first element is 1 (since there are no numbers to the left of the first element in nums).
		Then, a loop is used to fill in the rest of the left array,
		where each element at index i is the product of the number at the same index in nums and the previous element in left.
	*/
	left[0] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	/*
		The right array is similar to left, but it stores the product of all the numbers to the right of the current index in nums.
		It is initialized such that the last element is 1 (since there are no numbers to the right of the last element in nums).
		Then, a loop is used to fill in the rest of the right array in reverse order,
		where each element at index i is the product of the number at the next index in nums and the next element in right.
	*/
	right[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	/*
		Finally, the output array is filled in by multiplying the corresponding elements in the left and right arrays.
		This gives the product of all the numbers in the original array except the one at the current index, as required.
	*/
	for i := 0; i < n; i++ {
		output[i] = left[i] * right[i]
	}

	return output
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(nums)
	fmt.Println(productExceptSelf(nums))
}
