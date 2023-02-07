package main

import (
    "fmt"
)

func main()  {
    var arr [5] int
    arr[3] = 100
    //arr[5] = 200 //index out of range
    fmt.Println(arr)
    
    arr2 := make([]int, 5) 
    arr2[3] = 100
    arr2[5] = 100 //panic: runtime error: index out of range [5] with length 5
    arr2 = append(arr2, 66) //using appen the memory will increase dynamically
    fmt.Println(arr2)
    fmt.Println(arr2[3:])
}
