package main

import "fmt"

/*
933. Number of Recent Calls
You have a RecentCounter class which counts the number of recent requests within a certain time frame.

Implement the RecentCounter class:

- RecentCounter() Initializes the counter with zero recent requests.

- int ping(int t) Adds a new request at time t, where t represents some time in milliseconds, and
returns the number of requests that has happened in the past 3000 milliseconds (including the new request).
Specifically, return the number of requests that have happened in the inclusive range [t - 3000, t].

It is guaranteed that every call to ping uses a strictly larger value of t than the previous call.
*/

type RecentCounter struct {
	requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{requests: []int{}}
}

func (this *RecentCounter) Ping(t int) int {
	/*The space complexity of this code is O(1),
		which means it uses a constant amount of additional memory regardless of the input size.
		This is because the code only uses a few variables (startTime, count, and the loop index i)
		that do not depend on the size of the input. The this.requests slice is modified in-place and does not
		create any additional copies or data structures.

	  The time complexity of this code is O(n),
	  	where n is the number of elements in the this.requests slice.
		This is because the code iterates over the this.requests slice once,
		checking each element against the startTime value.
		The time it takes to execute the loop depends on the number of elements in the slice.
	*/

	this.requests = append(this.requests, t)
	startTime := t - 3000
	count := 0
	for i := len(this.requests) - 1; i >= 0; i-- {
		if this.requests[i] >= startTime {
			count++
		} else {
			break
		}
	}
	return count
}

func main() {
	counter := Constructor()
	fmt.Println(counter.Ping(1000)) // Output: 1
	fmt.Println(counter.Ping(2000)) // Output: 2
	fmt.Println(counter.Ping(3001)) // Output: 3
	fmt.Println(counter.Ping(3002)) // Output: 3
}
