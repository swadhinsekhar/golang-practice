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
    lock        *sync.RWMutex
    wg          *sync.WaitGroup
    readChan    chan string
    writeChan   chan string
    exitChan    chan bool
    counter     int32
}

func worker1(conn *connWorker) {
    done := false
    connIdleTicker := time.NewTicker(time.Second * time.Duration(10))
    for !done {
        select {
        case <- connIdleTicker.C: {
            fmt.Println("worker1 sending pong")
            conn.lock.Lock()
            conn.counter++
            conn.lock.Unlock()
            conn.exitChan <- true
            done = true
            fmt.Println("worker1 exiting")
        }
        case buf := <- conn.readChan: {
            fmt.Println("worker1 read buf: ", string(buf))
        }
        }
    }
    conn.wg.Done()
}

func worker2(conn *connWorker) {
    done := false
    connIdleTicker := time.NewTicker(time.Second * time.Duration(3))
    for !done {
        select {
        case <- connIdleTicker.C: {
            fmt.Println("worker2 sending pong")
            conn.lock.Lock()
            conn.counter++
            conn.lock.Unlock()
            buf := "hello from worker2"
            conn.readChan <- buf
        }
        case <- conn.exitChan:
            fmt.Println("worker2 exiting")
            done = true
        }
    }
    conn.wg.Done()
}

func main()  {
    fmt.Println("examples of go routine")

    core := make(map[string] * connWorker)

    conn := &connWorker {
        wg: &sync.WaitGroup{},
        lock: &sync.RWMutex{},
        readChan: make(chan string, 100),
        writeChan: make(chan string, 100),
        exitChan: make(chan bool, 1),
    }
    core["w1"] = conn
    core["w2"] = conn

    conn.wg.Add(1)
    go worker1(conn)

    conn.wg.Add(1)
    go worker2(conn)

    conn.wg.Wait()
    fmt.Println(conn.counter)

    close(conn.readChan)
    close(conn.writeChan)
    close(conn.exitChan)

    delete(core, "w1")
    delete(core, "w2")

    fmt.Println("main exiting...")
}
