package main

import "fmt"

func Max(x int, y int) int {
    if ( x > y) {
        return x
    }
    return y
}

//n : subarray
func maxSubarraySum(array [] int, n int) int {

    if (n > len(array)) {
        return 0
    }
    maxSum := 0 //holds the max sum of array
    tempSum := 0 //hold the sum of current subarray

    //sum of the first set of numbers
    for i := 0; i < n; i++ {
        maxSum += array[i]
    }
    tempSum = maxSum

    for i := n; i < len(array); i++ {
        /*
			assigning new tempsum,
			we get the current tempSum,
			move down the number of indexes == to n,
			subtract that number, then add the number of
			the current loops index
		*/
        //fmt.Println(tempSum, array[i-n], array[i])

        tempSum = tempSum - array[i-n] + array[i]
        maxSum = Max(maxSum, tempSum)
    }

    return maxSum
}

func main() {
    //fmt.Println(maxSubarraySum([]int{-1, -1, -2, 4, 2, 3, 5, 1}, 4))
	//fmt.Println(maxSubarraySum([]int{}, 4))
	fmt.Println(maxSubarraySum([]int{100, 200, 300, 400}, 2))
}
