package gobyexample

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// MutexesDemo - demonstrates using mutexes to safely access data across
// mutiple goroutines
func MutexesDemo() {
	// manage our state in a map
	var state = make(map[int]int)

	// this mutex will synchronize access to `state`
	var mutex = &sync.Mutex{}

	// keep track of how many read and write operations we do
	var readOps uint64
	var writeOps uint64

	// start 100 goroutines to execute repeated reads against the
	// state, once per millisecond in each goroutine
	for r := 0; r < 100; r++ {
		go func() {
			// for each read, we pick a key to access, Lock() the mutex
			// to ensure exclusive access to the `state`, read the value
			// at the chosen key, Unlock() the mutex, and increment the
			// `readOps` count
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				// wait a bit between reads
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// we also start 10 goroutines to simulate writes, using the same
	// pattern we did for reads
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// let the 10 goroutines work on the `state` and `mutex` for a second
	time.Sleep(time.Second)

	// take and report final operation counts
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// with a final lock of `state`, show how it ended up
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
