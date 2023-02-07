package main

import (
    "fmt"
    "sync"
    "time"

)

/*
    1. create two go routine should run independently
    2. add sync
    3. add various channel onexit, onread, onwrite
    4. should generate a ping pong in every 30 secs
*/

type connWorker struct {
    wg          *sync.WaitGroup
    readChan    chan *string
    writeChan   chan *string
    exitChan    chan bool

}

func worker1(conn *connWorker) {
    done := false
    connIdleTicker := time.NewTicker(time.Second * time.Duration(5))
    for !done {
        select {
        case <- connIdleTicker.C:
            fmt.Println("worker1 sending pong")
        }
    }
    conn.wg.Done()
}

func worker2(conn *connWorker) {
    done := false
    connIdleTicker := time.NewTicker(time.Second * time.Duration(3))
    for !done {
        select {
        case <- connIdleTicker.C:
            fmt.Println("worker2 sending pong")
    
        }
    }
    conn.wg.Done()
}

func main()  {
    fmt.Println("examples of go routine")
    conn := &connWorker {
        wg: &sync.WaitGroup{},
        readChan: make(chan *string, 100),
        writeChan: make(chan *string, 100),
        exitChan: make(chan bool, 1),
    }

    conn.wg.Add(1)
    go worker1(conn)

    conn.wg.Add(1)
    go worker2(conn)

    conn.wg.Wait()
   
    fmt.Println("exiting...")
}
