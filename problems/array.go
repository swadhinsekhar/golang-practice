package main

import (
	"fmt"
	"sort"
)

func main() {
	/*known length*/
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr1)
	fmt.Println(arr2)

	/*inferred lengths*/
	var arr3 = [...]int{1, 2, 3}
	arr4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)
	fmt.Println(arr4)

	/*string of arrays*/
	cars := [...]string{"volvo", "bwm", "suzuki", "lambogrini"}
	fmt.Println(cars)
	fmt.Println(cars[0])
	fmt.Println(cars[2])
	fmt.Println(cars[3])

	/*Array manumulation*/
	prices := [...]int{10, 20, 30}
	fmt.Println(prices)
	prices[1] = 50
	fmt.Println(prices)

	/*Array Initilaization*/
	var ar = [5]int{}               //not initialize
	var ar1 = [5]int{1, 2, 3}       //partially initialize
	var ar2 = [5]int{1, 2, 3, 4, 5} //fully initialize
	fmt.Println(ar)
	fmt.Println(ar1)
	fmt.Println(ar2)

	/*specific index initialize*/
	ar3 := [5]int{1: 50, 4: 60}
	fmt.Println(ar3)

	ar4 := [5]string{"volvo", "bwm", "suzuki", "lambogrini"}
	ar5 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(len(ar4))
	fmt.Println(len(ar5))

	var brr = []int{}
	fmt.Printf("brr: %v\n", brr)
	brr = append(brr, 3, 2, 1)
	fmt.Printf("brr: %v\n", brr)
	sort.Ints(brr)
	fmt.Printf("brr: %v\n", brr)

	for v := range brr {
		fmt.Printf("%v\n", v)
	}
}
