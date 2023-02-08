package main

import (
	"fmt"
	"sync"
)

/*
sysout 1 - 100
	1 - goroutine
		- odd numbers
		- 1, 3, 5, 7
	2 - goroutine
		- even numbers
		- 2, 4, 6, 8
	stop at 100

	communication between goroutine
*/

type sysout struct {
    oddChan chan int
    evenChan chan int
    exitOdd chan bool
}

func oddRoutine(wg *sync.WaitGroup, s *sysout) {
    done := false
    for !done {
        select {
        case oddData := <- s.oddChan:
            fmt.Println(oddData)
            s.evenChan <- oddData + 1
        case <- s.exitOdd:
            done = true
        }
    }
    wg.Done()
}

func evenRoutine(wg *sync.WaitGroup, s *sysout) {
    done := false
    for !done {
        select {
        case evenData := <- s.evenChan:
            fmt.Println(evenData)
            if (evenData == 100) {
                done = true
                s.exitOdd <- true
            }
            s.oddChan <- evenData + 1
        }
    }
    wg.Done()
}

func main() {
    fmt.Println("2 * goroutine print 1 to 100 synchronizingly ")
    
    s := &sysout {
        oddChan: make(chan int, 100),
        evenChan: make(chan int, 100),
        exitOdd: make(chan bool, 1),
    }
    wg := &sync.WaitGroup{}

    wg.Add(2)

    go oddRoutine(wg, s)
    go evenRoutine(wg, s)

    s.oddChan <- int(1)
    wg.Wait()
}
