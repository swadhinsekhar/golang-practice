package main

import (
    "fmt"
    . "pkg-test/pkg/utils"
    . "pkg-test/pkg/structs"
)


func main() {
    gc := &GlobalConfig{}
    gc.Url  = "192.168.1.100"
    fmt.Println("Hello FromMain")
    FromUtils()
    fmt.Println(gc.Url)
}
