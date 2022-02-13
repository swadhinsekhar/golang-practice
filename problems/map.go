package main

import (
    "fmt"
)


func main() {

    var a = make(map[string] string)

    a["brand"] = "ford"
    a["model"] = "Mustang"
    a["year"] = "1940"

    fmt.Printf("a: %v\n", a)
    fmt.Printf("a[branc]: %v\n", a["brand"])

    a["year"] = "1970" //modified
    a["color"] = "red" //added
    fmt.Printf("a after modified/added : %v\n", a)

    delete(a, "year")
    fmt.Printf("a after delete: %v\n", a)

    val1, ok1 := a["brand"]
    fmt.Printf("val1: %v ok1: %v\n", val1, ok1) // ford, true

    val2, ok2 := a["year"]
    fmt.Printf("val2: %v ok2: %v\n", val2, ok2) //empty string, false

    a["day"] = ""
    val3, ok3 := a["day"]
    fmt.Printf("val3: %v ok3: %v\n", val3, ok3) //empty string, true

    _, ok4 := a["model"]
    fmt.Printf("ok4: %v\n", ok4) //true

    var b = make(map[string] int)
    var c  map[string] int

    fmt.Printf("b: %v\n", (b == nil)) //false - memory created
    fmt.Printf("c: %v\n", (c == nil)) //true - no memory created


    var num = map[string] int {"one" : 1, "two": 2, "three" : 3, "four" : 4}

    for k, v := range num {
        fmt.Printf("[%v : %v]", k, v) //unsorted print
    }
    fmt.Println()

    var list = []string {}
    list = append(list, "one", "two", "three", "four")

    for _, v := range(list) {
        fmt.Printf("[%v : %v]", v, num[v])
    }
    fmt.Println()






}
