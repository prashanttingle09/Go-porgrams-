package main

import "fmt"

func main() {
	arr := []int{1, 5, 8, 7, 6, 2}
	channelData := Generator(arr)
	finalData := Squere(channelData)
	FinalPrint(finalData)
}

func Generator(arr []int) <-chan int {
	out := make(chan int)

	go func() {
		for data := range arr {
			out <- data
		}
		close(out)
	}()
	return out
}

func Squere(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for data := range in {
			out <- data * data
		}
		close(out)
	}()
	return out
}

func FinalPrint(in <-chan int) {
	for data := range in {
		fmt.Println("final data comes:", data)
	}
}
