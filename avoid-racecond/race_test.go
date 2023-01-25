package main

import (
	"sync/atomic"
	"testing"
)

func TestDataRaceCondition(t *testing.T) {
/*
    var counter int32
    for i := 0; i < 10; i++ {
        go func (i int) {
            counter += int32(i)
        }(i)
    }
*/
    var counter int32
    for i := 0; i < 10; i++ {
        go func (i int) {
            atomic.AddInt32(&counter, int32(i))
        }(i)
    }
}
