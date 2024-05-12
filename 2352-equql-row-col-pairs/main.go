package main

import "fmt"

/*
2352. Equal Row and Column Pairs

Given a 0-indexed n x n integer matrix grid, return the number of pairs (ri, cj) such that row ri and column cj are equal.

A row and column pair is considered equal if they contain the same elements in the same order (i.e., an equal array).

Example 1:
Input: grid = [[3,2,1],[1,7,6],[2,7,7]]
Output: 1
Explanation: There is 1 equal row and column pair:
- (Row 2, Column 1): [2,7,7]

Example 2:
Input: grid = [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]
Output: 3
Explanation: There are 3 equal row and column pairs:
- (Row 0, Column 0): [3,1,2,2]
- (Row 2, Column 2): [2,4,2,2]
- (Row 3, Column 2): [2,4,2,2]
*/

const ARR_LENGTH = 200

func equalPairs(grid [][]int) int {
	n := len(grid)
	rows := make(map[[ARR_LENGTH]int]int)
	cols := make(map[[ARR_LENGTH]int]int)
	count := 0

	for i := 0; i < n; i++ {
		row := [ARR_LENGTH]int{}
		col := [ARR_LENGTH]int{}

		for j := 0; j < n; j++ {
			row[j] = grid[i][j]
			col[j] = grid[j][i]
		}

		rows[row]++
		cols[col]++
	}

	for k, v := range rows {
		if cols[k] > 0 {
			count += cols[k] * v
		}
	}
	return count
}

func main() {

	grid := [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}

	fmt.Println(equalPairs(grid))

	grid = [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}
	fmt.Println(equalPairs(grid))
}
