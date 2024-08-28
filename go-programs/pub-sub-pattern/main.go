package main

import (
	"fmt"
	"sync"
)

func main() {
	topic := NewTopic()
	topic.wg.Add(2)
	go topic.Publisher()
	go topic.Subscriber()
	topic.wg.Wait()
}

type Message struct {
	Content string
}

type Topic struct {
	wg sync.WaitGroup
	ch chan Message
}

func NewTopic() *Topic {
	return &Topic{ch: make(chan Message, 5)}
}

func (t *Topic) Publisher() {
	defer t.wg.Done()
	for i := 1; i <= 5; i++ {
		msg := Message{Content: fmt.Sprintf("Message content:%v", i)}
		t.ch <- msg
	}
	close(t.ch)
}

func (t *Topic) Subscriber() {
	defer t.wg.Done()
	for data := range t.ch {
		fmt.Println("Message recived:", data)
	}
}

func (t *Topic) Subscribe() chan Message {
	return t.ch
}
