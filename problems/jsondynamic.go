/*
{
    "name": "gaurav",
    "age": 20,
    "address": "bangalore",
    "height": "5.8"
}

Output -
Name is string
Age is integer
Address is string
Height is float

JSON body where the keys are always string,
but the number of keys and keys themselves are not static.
The values can be either integer, string or float.

Write a program which can take this as a byte array, convert
it to the required data structure and print the value and the type
of the value

*/

package main

import (
    "fmt"
)

func translate(body []byte) {
    //map
    //refelect
    //typeof
    //inside map cast
    fmt.Println(body)
}

func main() {
    body := [] byte ("\"name\": \"gaurav\", \"age\": 20, \"address\": \"bangalore\"")

    translate(body)
}







