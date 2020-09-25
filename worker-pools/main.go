package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d has started job %d\n", id, j)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Printf("Worker %d has finished job %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	const NUMJOBS = 5
	jobs := make(chan int, NUMJOBS)
	results := make(chan int, NUMJOBS)
	defer close(results)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= NUMJOBS; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= NUMJOBS; a++ {
		fmt.Println(<-results)
	}
}
