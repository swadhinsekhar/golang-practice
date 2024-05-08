package main

import "fmt"

/*
605. Can Place Flowers

You have a long flowerbed in which some of the plots are planted, and some are not.
However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty,
 and an integer n, return true if n new flowers can be planted in the flowerbed without
 violating the no-adjacent-flowers rule and false otherwise.

Example 1:

Input: flowerbed = [1,0,0,0,1], n = 1
Output: true
Example 2:

Input: flowerbed = [1,0,0,0,1], n = 2
Output: false

*/

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i := 0; i < len(flowerbed); i++ {
		/*
			it checks if the plot is empty (i.e., flowerbed[i] == 0) and
			if the plots adjacent to it are also empty or if the plot is at the edge of the
			flowerbed (i.e., i == 0 || flowerbed[i-1] == 0 for the plot before and i == len(flowerbed)-1 || flowerbed[i+1] == 0 for the plot after).
			If all these conditions are met, it means a flower can be planted at this plot without violating the rule. So,
			it plants a flower (i.e., sets flowerbed[i] to 1) and increments the count by 1.
		*/
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) &&
			(i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			count++
		}
	}
	return count >= n
}

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
}
