package designpatterns

import (
	"fmt"
)

type Message struct {
	Content string
}

type Topic struct {
	ch chan Message
}

func NewTopic() *Topic {
	return &Topic{ch: make(chan Message, 5)}
}
func (t *Topic) Subscbe() chan Message {
	return t.ch
}
func Publisher(t *Topic) {
	for i := 1; i < 6; i++ {
		cnt := fmt.Sprintf("Publisher Publish: %v", i)
		msg := Message{Content: cnt}
		t.ch <- msg
	}
	close(t.ch)
}

func Subscriber(ch chan Message) {
	for data := range ch {
		fmt.Println("Data Recived:", data)
	}
}

/*
func main() {
	topic := NewTopic()
	go Publisher(topic)
	go Subscriber(topic.Subscbe())
	fmt.Println("Started:")
	time.Sleep(5 * time.Second)
}
*/
