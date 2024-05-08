package main

import (
	"fmt"
)

func main() {
	var arr [5]int
	arr[3] = 100
	//arr[5] = 200 //index out of range
	fmt.Println(arr)

	arr2 := make([]int, 5)
	arr2[3] = 100
	//arr2[5] = 100           //panic: runtime error: index out of range [5] with length 5
	arr2 = append(arr2, 66) //using appen the memory will increase dynamically
	fmt.Println(arr2)
	fmt.Println(arr2[3:])

	/*
		- The `copy` function is used to copy elements between slices or arrays.
		- Here's an example of copying elements from one slice to another:
	*/
	src := []int{1, 2, 3, 4, 5}
	dest := make([]int, len(src))
	//copy(dest, src)
	n := copy(dest, src)
	fmt.Println("==> array-to-slice copy ", dest, "num of elements copied: ", n)

	/*
		- Slices can be sliced to create new slices with different lengths and capacities.
		- Here's an example of slicing a slice:
	*/
	slice := []int{1, 2, 3, 4, 5}
	newSlice := slice[0:3] // 0, 1, 2 : 3 is exclusive, like array index starts from 0
	fmt.Println("==> new slice ", newSlice)

	/*
	   - converting array to slice
	*/
	arr3 := [5]int{1, 2, 3, 4, 5}
	slice2 := arr3[:]
	fmt.Println("==> array to slice ", slice2)

	/*
		- Slices have a length and capacity.
		- The length of a slice is the number of elements it contains.
		- The capacity of a slice is the number of elements in the underlying array, starting from the index where the slice begins.
		- The `append` function is used to add elements to a slice. If the length exceeds the capacity, a new underlying array is allocated and the elements are copied.
		- Here's an example of appending elements to a slice:
	*/
	/*
	  output:
	          ==> slice4 before append  [1 2 3 4 5] length:  5 capacity:  5
	          ==> slice4 after append  [1 2 3 4 5 6] length:  6 capacity:  10
	  why capacity is changed from 5 to 10?
	  - Go checks if there's enough capacity in the underlying array to accommodate
	  the new element (6). Since the current capacity is 5 and not sufficient,
	  Go creates a new underlying array with double the capacity (10 in this case).
	  - Go then copies the existing elements from the old array to the new array and appends
	  the new element (6) to the end. After appending, slice4 now refers to the new array,
	  which has a capacity of 10, and its length is updated to 6.
	  -  it's essential to keep in mind that occasionally, this behavior may lead to
	  more memory being allocated than actually required.
	*/
	slice4 := []int{1, 2, 3, 4, 5}
	fmt.Println("==> slice4 before append ", slice4, "length: ", len(slice4), "capacity: ", cap(slice4))
	slice4 = append(slice4, 6)
	fmt.Println("==> slice4 after append ", slice4, "length: ", len(slice4), "capacity: ", cap(slice4))
}
