# Race condition in golang

### instead of sync.Mutex use "sync/atomic" package to avoid race condition

```
    var counter int32
    for i := 0; i < 10; i++ {
        go func (i int) {
            counter += int32(i)
        }(i)
    }
```

### using atomic adding of int32

```
    var counter int32
    for i := 0; i < 10; i++ {
        go func (i int) {
            atomic.AddInt32(&counter, int32(i))
        }(i)
    }

```
