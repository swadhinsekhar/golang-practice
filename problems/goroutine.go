package main

import (
    "fmt"
    "sync"
    "os"
    "os/signal"
    "syscall"
    "time"
)

type pingReply struct {
    Status              string
}

type spsc struct {
    mu                  sync.Mutex
    wg                  *sync.WaitGroup
    pingConsumerChan    chan * pingReply
    pingProducerChan    chan * pingReply
    sigChan             chan os.Signal
    exitConsRoutine     chan bool
    exitProdRoutine     chan bool
}

func (s *spsc) producer() {

    defer s.wg.Done()

    pingTicker := time.NewTicker(time.Second * time.Duration(1))

    done := false
    for; !done; {
        select {
        case <-pingTicker.C:
            fmt.Printf("Ping sent to Consumer Chan\n")
            s.pingConsumerChan <- &pingReply{Status: "success"}

        case <-s.pingProducerChan:
            fmt.Printf("Ping Receive from Consumer Chan\n")

        case <-s.exitProdRoutine:
            fmt.Printf("Producer is exiting..\n")
            done = true
        }
    }
}

func (s *spsc) consumer() {

    defer s.wg.Done()

    pingTicker := time.NewTicker(time.Second * time.Duration(1))

    done := false
    for; !done; {
        select {
        case <-pingTicker.C:
            fmt.Printf("Ping sent to Producer Chan\n")
            s.pingProducerChan <- &pingReply{Status: "success"}

        case <-s.pingConsumerChan:
            fmt.Printf("Ping Receive from Producer Chan\n")

        case <-s.exitConsRoutine:
            fmt.Printf("Consumer is exiting..\n")
            done = true
        }
    }
}

func main() {

    s := &spsc{}

    s.wg = &sync.WaitGroup{}
    s.pingConsumerChan = make(chan * pingReply, 1)
    s.pingProducerChan = make(chan * pingReply, 1)
    s.exitConsRoutine = make(chan bool, 1)
    s.exitProdRoutine = make(chan bool, 1)
    s.sigChan = make(chan os.Signal, 2)
    signal.Notify(s.sigChan, syscall.SIGINT, syscall.SIGTERM)

    s.wg.Add(2)
    go s.producer()
    go s.consumer()

    done := false
    for ; !done; {
        select {

        case <-s.sigChan:
            s.exitProdRoutine <- true
            s.exitConsRoutine <- true
            fmt.Printf("Exiting process started from main\n")
            close(s.pingConsumerChan)
            close(s.pingProducerChan)
            time.Sleep(2 * time.Second)
            done = true

        default:
            time.Sleep(1 * time.Second)
        }
    }

    s.wg.Wait()
    os.Exit(1)
}
