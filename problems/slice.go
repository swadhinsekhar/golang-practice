package main

import "fmt"

/*
    Slices are similar to arrays, but are more powerful and flexible.

    Like arrays, slices are also used to store multiple values of the same type in a single variable.

    However, unlike arrays, the length of a slice can grow and shrink as you see fit.

    In Go, there are several ways to create a slice:

        Using the []datatype{values} format
        Create a slice from an array
        Using the make() function
*/


func main() {
    /*
        Syntax for declaring slice:

        slice_name := []datatype{values}
        myslice := []int
        myslice := []int{1,2,3}

        len() function - returns the length of the slice (the number of elements in the slice)
        cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
    */

    myslice := []int{}

    fmt.Println(len(myslice))
    fmt.Println(cap(myslice))
    fmt.Println(myslice)

    myslice2 := [10] string {"Go", "Slices", "are", "powerfull"}
    fmt.Println(len(myslice2))
    fmt.Println(cap(myslice2))
    fmt.Println(myslice2)

    /*Create a slice from an array*/
    arr := [6] int {10, 20, 30, 40, 50, 60}
    myslice3 := arr[2:4]
    fmt.Printf("len : %d\n", len(myslice3))
    fmt.Printf("cap : %d\n", cap(myslice3))
    fmt.Printf("values : %v\n", (myslice3))



    /* Using the make() function */
    //slice_name := make([]type, length, capacity)

    myslice4 := make([]int, 5, 10)
    fmt.Printf("myslice4 len : %d\n", len(myslice4))
    fmt.Printf("myslice4 cap : %d\n", cap(myslice4))
    fmt.Printf("myslice4 values : %v\n", (myslice4))

    myslice5 := make([]int, 5)
    fmt.Printf("myslice5 len : %d\n", len(myslice5))
    fmt.Printf("myslice5 cap : %d\n", cap(myslice5))
    fmt.Printf("myslice5 values : %v\n", (myslice5))


    //access/modify of slice
    prices := [] int {10, 20, 30}
    fmt.Printf("0th>> %v\n", prices[0])
    fmt.Printf("2nd>> %v\n", prices[2])
    //fmt.Printf("3rd>> %v\n", prices[3]) //error index out of range
    prices[2] = 50
    fmt.Printf("0th>> %v\n", prices[0])
    fmt.Printf("2nd>> %v\n", prices[2])


    //append elements to the slices
    numslice := [] int {1, 2, 3, 4, 5}
    fmt.Printf("numslice values : %v\n", (numslice))
    fmt.Printf("numslice len : %d\n", len(numslice))
    fmt.Printf("numslice cap : %d\n", cap(numslice))
    numslice = append(numslice, 10, 20, 30, 40, 50)
    fmt.Printf("numslice values : %v\n", (numslice))
    fmt.Printf("numslice len : %d\n", len(numslice))
    fmt.Printf("numslice cap : %d\n", cap(numslice))

    //append two slices to third slice
    myse1 := [] int {1, 2, 3}
    myse2 := [] int {4, 5, 6}
    myse3 := append(myse1, myse2...)
    fmt.Printf("myse3 value: %v\n", myse3)
    fmt.Printf("myse3 len: %v\n", len(myse3))
    fmt.Printf("myse3 cap: %v\n", cap(myse3))


    //change the length of a slice:
    arr1 := [6] int {1,2,3,4,5,6}
    myslice6 := arr1[1:5] //slice array
    fmt.Printf("myslice6 value: %v\n", myslice6)
    fmt.Printf("myslice6 len: %v\n", len(myslice6))
    fmt.Printf("myslice6 cap: %v\n", cap(myslice6))

    myslice6 = arr1[1:3] //change of the length by re-slice
    fmt.Printf("myslice6 value: %v\n", myslice6)
    fmt.Printf("myslice6 len: %v\n", len(myslice6))
    fmt.Printf("myslice6 cap: %v\n", cap(myslice6))

    myslice6 = append(myslice6, 20, 22, 23, 25) //change the length by appending items
    fmt.Printf("myslice6 value: %v\n", myslice6)
    fmt.Printf("myslice6 len: %v\n", len(myslice6))
    fmt.Printf("myslice6 cap: %v\n", cap(myslice6))

    numbers := [...] int {1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
    fmt.Printf("numbers value: %v\n", numbers)
    fmt.Printf("numbers len: %v\n", len(numbers))
    fmt.Printf("numbers cap: %v\n", cap(numbers))

    neededNumbers := numbers[:len(numbers) - 10]
    numbersCopy := make([] int , len(neededNumbers))
    copy(numbersCopy, neededNumbers)

    fmt.Printf("numbersCopy: %v\n", numbersCopy)
    fmt.Printf("numbersCopy len: %v\n", len(numbersCopy))
    fmt.Printf("numbersCopy cap: %v\n", cap(numbersCopy))

    for idx, val := range numbersCopy {
        fmt.Printf("%v : %v\n", idx, val)
    }

}
