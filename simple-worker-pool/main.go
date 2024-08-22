package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	numsJob := 10
	numsWorker := 3
	jobs := make(chan int, numsJob)
	result := make(chan int, numsJob)

	for i := 1; i <= numsWorker; i++ {
		wg.Add(1)
		go worker(jobs, result, &wg, i)
	}

	for j := 1; j <= numsJob; j++ {
		jobs <- j
	}

	close(jobs)
	wg.Wait()
	close(result)
	for c := 1; c <= numsJob; c++ {
		fmt.Printf("result collected:%d \n", <-result)
	}
}

func worker(jobs chan int, result chan int, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for res := range jobs {
		fmt.Printf("worker %d started to work on :%d \n", id, res)
		result <- res * 2
		// runtime.Gosched()
	}
}
