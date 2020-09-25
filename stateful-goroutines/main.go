package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type ReadOp struct {
	key  int
	resp chan int
}

type WriteOp struct {
	key, val int
	resp     chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan ReadOp)
	writes := make(chan WriteOp)

	// state is owned by this goroutine
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// writers
	for r := 0; r < 1000; r++ {
		go func() {
			for {
				read := ReadOp{key: rand.Intn(5), resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// writers
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := WriteOp{key: rand.Intn(5), val: rand.Intn(100), resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Printf("readOpsFinal = %d, writeOpsFinal = %d\n", readOpsFinal, writeOpsFinal)
}
