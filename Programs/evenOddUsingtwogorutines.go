// package main

// import (
// 	"fmt"
// 	"time"
// )

// func printEven(ch chan bool) {
// 	for i := 2; i <= 10; i = i + 2 {
// 		<-ch
// 		fmt.Println("even goroutine:", i)
// 		ch <- true
// 	}
// }

// func printOdd(ch chan bool) {
// 	for i := 1; i <= 10; i = i + 2 {
// 		fmt.Println("odd goroutine:", i)
// 		ch <- true
// 		<-ch
// 	}
// }

// func main() {
// 	ch := make(chan bool)

// 	go printEven(ch)
// 	go printOdd(ch)
// 	time.Sleep(5 * time.Second)
// 	close(ch)
// }
