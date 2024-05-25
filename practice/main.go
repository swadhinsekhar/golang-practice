package main

import "fmt"

func rotatedArray(arr []int) []int {
	l := len(arr) - 1
	for i, j := 0, l; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println(rotatedArray(arr))
}
