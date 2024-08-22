package main

import (
	"fmt"
	"sync"
)

// print even and odd nuber using 2 gorutines in sequential order
func main() {
	odch := make(chan bool)
	evenCh := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go PrintOdd(evenCh, odch, &wg)
	go PrintEven(evenCh, odch, &wg)
	odch <- true
	wg.Wait()
}

func PrintEven(evenCh chan bool, oddCh chan bool, wg *sync.WaitGroup) {
	n := 100
	for i := 2; i <= n; i = i + 2 {
		<-evenCh
		fmt.Println("Even Gouroutine:", i)
		if i < n {
			oddCh <- true

		}
	}

	wg.Done()
}

func PrintOdd(evenCh chan bool, oddCh chan bool, wg *sync.WaitGroup) {
	n := 100
	for i := 1; i <= n; i = i + 2 {
		<-oddCh
		fmt.Println("Odd Gouroutine:", i)
		evenCh <- true
	}
	wg.Done()
}
