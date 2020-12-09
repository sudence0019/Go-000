package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var count1 int32 = 0
var aountAtomic atomic.Value
var wg sync.WaitGroup
var mutex sync.Mutex

func addCountByDefault() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				count1++
				time.Sleep(2 * time.Millisecond)
			}
			defer func() { wg.Done() }()
		}()
	}
}

func addCountByMutex() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				mutex.Lock()
				count1++
				mutex.Unlock()
				time.Sleep(2 * time.Millisecond)
			}
			defer func() { wg.Done() }()
		}()
	}
}

func addCountByAtomic() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {

				aountAtomic.Store(
					atomic.AddInt32(&count1, 1))

				time.Sleep(2 * time.Millisecond)
			}
			defer func() { wg.Done() }()
		}()
	}
}
func main() {
	addCountByAtomic()
	wg.Wait()
	fmt.Println(aountAtomic.Load())
}
