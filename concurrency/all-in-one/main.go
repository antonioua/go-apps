package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// Goroutine
	go func() {
		fmt.Println("Spawned new goroutine.")
	}()
	fmt.Println("This is main goroutine.")

	// channel
	ch := make(chan string)
	go func() {
		ch <- "Hello from goroutine!"
	}()
	fmt.Println(<-ch)

	// Buffered channel
	// When we need to send more that one message to channel at a time
	ch = make(chan string, 2)
	go func() {
		ch <- "Hello from goroutine message1!"
		ch <- "Hello from goroutine message2!"
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Wait group
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Inside goroutine.")
	}()
	wg.Wait()
	fmt.Println("This is main goroutine.")

	// Mutex
	var mu sync.Mutex
	wg = sync.WaitGroup{}

	sum := 0
	iterations := 1000

	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			sum++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Sum: ", sum)

	// Once
	var once sync.Once
	sum = 0

	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			defer wg.Done()
			once.Do(func() {
				sum++
			})
		}()
	}
	wg.Wait()
	fmt.Println("Sum: ", sum)

	// Pool
	memPool := &sync.Pool{
		New: func() interface{} {
			mem := make([]byte, 1024)
			return &mem
		},
	}
	mem := memPool.Get().(*[]byte) // Cast to pointer of slice of bytes
	// Here should go some instructions to utilize such piece of memory
	memPool.Put(mem) // Put back the memory to the pool so we avoid performing additional allocation

	// Cond
	// Used to signal or broadcast between goroutines
	c := sync.NewCond(&sync.Mutex{})
	go func() {
		c.L.Lock()
		// changing some condition
		c.L.Unlock()
		c.Signal()
	}()
	// Checking the condition ...
	c.L.Lock()
	c.Wait()
	c.L.Unlock()

	// Goroutine safe Map
	syncMap := sync.Map{}

	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			defer wg.Done()
			syncMap.Store(0, i)
		}()
	}
	wg.Wait()
	key, value := syncMap.Load(0)
	fmt.Println("Sync map.\nKey: ", key, "\nValue: ", value)

	// Atomic
	var i int64
	atomic.AddInt64(&i, 1)
	fmt.Println("Atomic: ", i)

	var av atomic.Value
	type ninja struct {
		name string
	}
	av.Store(ninja{name: "Ninja Golang"})
	fmt.Println(av.Load())
}
