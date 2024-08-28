package concurrency

import "time"

// Race condition  -happend when two gorotuines try to access or modify the shared resource can detect race using go run main.go --race

var counter int

func Race() {
	go increment()
	go increment()
	time.Sleep(1 * time.Second)
}

func increment() {
	for i := 0; i < 1000; i++ {
		counter++
	}
}
