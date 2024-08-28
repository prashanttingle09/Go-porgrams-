package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	numJob := 10
// 	jobs := make(chan string, numJob)
// 	result := make(chan string, numJob)

// 	for w := 0; w < 3; w++ {
// 		go Worker(jobs, result)
// 	}

// 	for i := 1; i <= numJob; i++ {
// 		jobs <- fmt.Sprintf("job:%v", i)
// 	}
// 	close(jobs)

// 	for i := 1; i <= numJob; i++ {
// 		fmt.Println("result:", <-result)
// 	}

// 	time.Sleep(5 * time.Second)
// }

// func Worker(jobs chan string, result chan string) {
// 	for data := range jobs {
// 		result <- data
// 	}
// }

// func main() {
// 	arr := []int{10, 20, 20, 30, 30, 30, 10}
// 	removeDuplicate(arr)
// 	fmt.Println("arre is:", arr)
// }

// func removeDuplicate(arr []int) {
// 	j := 1
// 	for i := 1; i < len(arr)-1; i++ {
// 		if arr[i] == arr[i+1] {
// 			arr[j] = arr[i]
// 			j++
// 		}
// 	}
// }

// func main() {
// 	constJob := 10
// 	workerNum := 3
// 	jobs := make(chan int, constJob)
// 	result := make(chan int, constJob)

// 	for i := 0; i < workerNum; i++ {
// 		go Worker(jobs, result)
// 	}

// 	for j := 1; j <= constJob; j++ {
// 		jobs <- j
// 	}
// 	close(jobs)

// 	for c := 1; c <= constJob; c++ {
// 		res := <-result
// 		if res != 0 {
// 			fmt.Println("result collected:", res)
// 		}
// 	}
// 	time.Sleep(5 * time.Second)
// }

// func Worker(jobs chan int, result chan int) {
// 	for job := range jobs {
// 		if isEven(job) {
// 			result <- job
// 		} else {
// 			result <- 0
// 		}
// 	}
// }

// func isEven(num int) bool {
// 	if num%2 == 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

type Message struct {
	Content string
}

type Topic struct {
	ch chan Message
}

func NewTopic() *Topic {
	return &Topic{ch: make(chan Message, 5)}
}
func main() {

}
