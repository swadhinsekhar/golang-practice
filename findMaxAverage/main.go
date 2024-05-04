package main

import (
	"fmt"
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	// Calculate the sum of the first k elements
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum

	// Slide the window and update the sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		maxSum = int(math.Max(float64(maxSum), float64(sum)))
	}

	// Calculate the average of the maximum sum
	return float64(maxSum) / float64(k)
}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Printf("Maximum average subarray of length %d: %.2f\n", k, findMaxAverage(nums, k))
}
