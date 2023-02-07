package main

import (
    "fmt"
    "math"
)

//interafce are collections of methods of signature
type geometry interface {
    area() float32
    perim() float32
}

type rect struct {
    width float32
    length float32
}

type circle struct {
    radius float32
}

func (c circle) area() float32 {
   return math.Pi * c.radius * c.radius 
}

func (c circle) perim() float32 {
   return 2 * math.Pi * c.radius
}

func (r rect) area() float32 {
   return r.width * r.length
}

func (r rect) perim() float32 {
   return 2 * r.length * 2 * r.width 
}

//method which have funcions
func measure(g geometry) {
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main()  {
    fmt.Println("interface are collections of methods of signature")

    mycirle := circle {
        radius: 5,
    }
    myrect := rect {
        width: 30,
        length: 20,
    }

    fmt.Println(mycirle.area())
    fmt.Println(mycirle.perim())
    fmt.Println(myrect.area())
    fmt.Println(myrect.perim())
    //instead of calling above function call by interface

    //if a method is missing then error will be : circle does not implement geometry (missing perim method)
    measure(mycirle)
    measure(myrect)

    
}
