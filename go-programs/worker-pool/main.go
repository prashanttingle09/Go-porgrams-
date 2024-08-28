package main

import (
	"fmt"
	"sync"
)

type Job struct {
	ID int
}
type workerpool struct {
	wg        sync.WaitGroup
	numWorker int
	result    chan int
	jobQueue  chan Job
}

func NewWorkerPool(numWorker int, jobQueueSize int) *workerpool {
	return &workerpool{
		numWorker: numWorker,
		jobQueue:  make(chan Job, jobQueueSize),
		result:    make(chan int, jobQueueSize),
	}
}

func (w *workerpool) worker(id int) {
	defer w.wg.Done()
	for job := range w.jobQueue {
		fmt.Printf("Worker %d started job %d\n", id, job.ID)
		fmt.Printf("Worker %d finished job %d\n", id, job.ID)
		w.result <- job.ID
	}
}

func (wp *workerpool) start() {
	for i := 1; i <= wp.numWorker; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}
func (ws *workerpool) wait() {
	ws.wg.Wait()
	close(ws.result)
}
func (ws *workerpool) sendJobs(job Job) {
	ws.jobQueue <- job
}

func (ws *workerpool) collectResult() {
	for result := range ws.result {
		fmt.Printf("Result received for job %d\n", result)
	}
}

func main() {
	numsWorker := 3
	numJobs := 10
	workerPool := NewWorkerPool(numsWorker, numJobs)
	for i := 1; i <= numJobs; i++ {
		workerPool.sendJobs(Job{ID: i})
	}
	close(workerPool.jobQueue)
	workerPool.start()
	workerPool.wait()
	workerPool.collectResult()
}
