package main

import (
	"fmt"
)

/*
* The Two Number Sum problem goes like this —
* you are given an array of integers called nums.
* You are also given an integer target, and
* your task is to see if there are two numbers in
* the array that add up to target. If there are,
* then return an array of the indices of the two numbers
* that add up to target. Otherwise, just return an empty array.
**/

func twoSum(array []int, target int) []int {
	for i := 0; i < len(array); i++ {

		var start = array[i]

		for j := i + 1; j < len(array); j++ {

			var end = array[j]

			if (start + end) == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSumAdv(array []int, target int) []int {
	seenNums := map[int]int{}
	for i, num := range array {
		potentialMatch := target - num
		if j, found := seenNums[potentialMatch]; found {
			return []int{i, j}
		}
		seenNums[num] = i
	}
	return []int{}
}

func main() {

	arr := []int{2, 7, 11, 15}
	fmt.Println(twoSum(arr, 9))
	fmt.Println(twoSum(arr, 13))
	fmt.Println(twoSumAdv(arr, 9))
	fmt.Println(twoSumAdv(arr, 13))
}

/*
func twoSums(array [] int, target int) [] int {
	for (i := 0; i < len(array); i++) {
		var start = array[i]
		for (j := i + 1; j < len(array); j++) {
			var end = array[j]
			if ((start + end) == target) {
				return [] int {i, j}
			}
		}
	}

	return [] int {}
}

func twoSumAdv(array [] int, target int) [] int {

	//arr := [] int {2,7,11,15}
	seenSums := map[int] int{}

	for i, num := range array {
		potentialNum := target - num
		if j, found := seenSums[potentialNum]; found {
			return [] int {i, j}
		}
		seenSums[num] = i
	}

	return [] int {}
}

*/
