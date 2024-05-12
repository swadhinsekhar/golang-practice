package main

import "fmt"

/*
	215. Kth Largest Element in an Array
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4

*/

func partition(nums []int, low, high int) int {
	// Choose the last element as the pivot
	pivot := nums[high]
	// Initialize the partition index as the lowest element
	partitionIndex := low
	// Iterate through the array
	for i := low; i < high; i++ {
		// If the current element is less than the pivot
		if nums[i] < pivot {
			// Swap the current element with the element at the partition index
			nums[i], nums[partitionIndex] = nums[partitionIndex], nums[i]
			// Increment the partition index
			partitionIndex++
		}
	}
	// Swap the pivot with the element at the partition index
	nums[partitionIndex], nums[high] = nums[high], nums[partitionIndex]
	// Return the partition index
	return partitionIndex
}

func quickSort(nums []int, low, high int) {
	stack := make([]int, high-low+1)
	top := -1

	// Push initial values of low and high to the stack
	top++
	stack[top] = low
	top++
	stack[top] = high

	// Keep popping from the stack until it is empty
	for top >= 0 {
		// Pop high and low
		high = stack[top]
		top--
		low = stack[top]
		top--

		// Find the partition index
		partitionIndex := partition(nums, low, high)

		// If there are elements on the left side of the partition,
		// push the low and high indices to the stack
		if partitionIndex-1 > low {
			top++
			stack[top] = low
			top++
			stack[top] = partitionIndex - 1
		}

		// If there are elements on the right side of the partition,
		// push the low and high indices to the stack
		if partitionIndex+1 < high {
			top++
			stack[top] = partitionIndex + 1
			top++
			stack[top] = high
		}
	}
}

func findKthLargest(nums []int, k int) int {
	// Sort the array in descending order
	quickSort(nums, 0, len(nums)-1)
	// Return the kth element
	return nums[k-1]
}

func main() {
	fmt.Println("215. Kth Largest Element in an Array!")
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(nums, k)
	fmt.Println("Output:", findKthLargest(nums, k))

	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k = 4
	fmt.Println(nums, k)
	fmt.Println("Output:", findKthLargest(nums, k))

}
