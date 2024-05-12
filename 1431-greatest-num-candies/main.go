package main

import "fmt"

/*
1431. Kids With the Greatest Number of Candies
There are n kids with candies. You are given an integer array candies, where each candies[i]
represents the number of candies the ith kid has, and an integer extraCandies, denoting the number of
extra candies that you have.

Return a boolean array result of length n, where result[i] is true if, after giving the ith kid all the
extraCandies, they will have the greatest number of candies among all the kids, or false otherwise.

Note that multiple kids can have the greatest number of candies.

Input: candies = [2,3,5,1,3], extraCandies = 3
Output: [true,true,true,false,true]
*/

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	// Find the max number of candies
	for _, c := range candies {
		if c > max {
			max = c
		}
	}

	result := make([]bool, len(candies))
	for i, c := range candies {
		// Check if the kid has the greatest number of candies
		/* For example, if c is 5, extraCandies is 3, and max is 7,
		* the expression c+extraCandies >= max will evaluate to true because
		* 5 + 3 is equal to 8, which is greater than or equal to 7.
		* */
		result[i] = c+extraCandies >= max
	}

	return result
}

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
}
