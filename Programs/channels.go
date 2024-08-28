package main

import "fmt"

func Send(ch chan string) {
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("Sending data:%v", i)
		ch <- msg
	}
	close(ch)
}

func Recive(ch chan string) {
	for data := range ch {
		fmt.Println("Data recivd in Recive gorutine:", data)
	}
}
