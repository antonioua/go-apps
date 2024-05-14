package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type sharedBuff struct {
	buffer []byte
	lock   sync.RWMutex
}

func (b *sharedBuff) Read() []byte {
	b.lock.RLock()
	defer b.lock.RUnlock()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	return b.buffer
}

func (b *sharedBuff) Write(data []byte) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.buffer = data
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
}

func main() {
	buff := &sharedBuff{}
	for {
		N := 5
		for i := 0; i <= N; i++ {
			go func(id int) {
				for {
					data := []byte(fmt.Sprintf("Data from goroutine writer #%d", i))
					buff.Write(data)
				}
			}(i)
		}

		M := 3
		for i := 0; i <= M; i++ {
			go func(id int) {
				for {
					data := buff.Read()
					fmt.Printf("\nReading data in goroutine #%d: %s\n", i, data)
				}
			}(i)
		}
	}
}
