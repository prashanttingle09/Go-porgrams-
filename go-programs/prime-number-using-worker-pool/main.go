package main

import (
	"fmt"
	"sync"
)

type Workerpool struct {
	wg          sync.WaitGroup
	JobsQueue   chan int
	Result      chan int
	WorkerCount int
}

func NewWorkerPool(jobsqueueCount int, workerCount int) *Workerpool {
	return &Workerpool{
		JobsQueue:   make(chan int, jobsqueueCount),
		Result:      make(chan int, jobsqueueCount),
		WorkerCount: workerCount,
	}
}

func main() {
	numsJobs := 20
	numsWorker := 3
	wp := NewWorkerPool(numsJobs, numsWorker)

	for w := 1; w <= numsWorker; w++ {
		wp.wg.Add(1)
		go wp.worker(w)
	}

	for j := 1; j <= numsJobs; j++ {
		wp.JobsQueue <- j
	}

	close(wp.JobsQueue)
	wp.wg.Wait()
	close(wp.Result)
	arr := []int{}

	for data := range wp.Result {
		if data != 0 {
			arr = append(arr, data)
		}
	}
	fmt.Println("All prime numbers from 1-20:", arr)

}

func (wp *Workerpool) worker(id int) {
	defer wp.wg.Done()
	for num := range wp.JobsQueue {
		fmt.Printf("worker: %d started to work on: %d\n", id, num)
		if isPrime(num) {
			wp.Result <- num
		} else {
			wp.Result <- 0
		}
	}
}

func isPrime(num int) bool {
	if num == 1 || num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}

	for i := 3; i < num; i = i + 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
