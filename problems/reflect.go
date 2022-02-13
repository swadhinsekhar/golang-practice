package main

import (
    "fmt"
    "reflect"
)

func parseInterface(evArgs interface{}) {

        switch reflect.TypeOf(evArgs).Kind() {
        case reflect.Slice:
            s := reflect.ValueOf(evArgs)
            for i := 0; i < s.Len(); i++ {
                fmt.Printf("%v", s.Index(i))
            }

        case reflect.String:
            s := reflect.ValueOf(evArgs)
            fmt.Printf("%s\n", s.String())

        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            s := reflect.ValueOf(evArgs)
            fmt.Printf("%v\n", s)
        }
}


func main() {
    parseInterface("hello")
    var val int = 55
    parseInterface(val)
}
