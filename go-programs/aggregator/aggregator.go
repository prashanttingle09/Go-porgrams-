package main

import (
	"fmt"
	"sync"
)

type Pool struct {
	SquereCh    chan int
	MultiplyeCh chan int
	wg          sync.WaitGroup
}

func NewPool() *Pool {
	return &Pool{
		SquereCh:    make(chan int),
		MultiplyeCh: make(chan int),
	}
}

func (p *Pool) squere(n int) {
	defer p.wg.Done()
	p.SquereCh <- n * n
	close(p.SquereCh)
}

func (p *Pool) multiply(n int) {
	defer p.wg.Done()
	p.MultiplyeCh <- n * n * n
	close(p.MultiplyeCh)
}

func (p *Pool) print() {
	defer p.wg.Done()
	for p.SquereCh != nil || p.MultiplyeCh != nil {
		select {
		case data, ok := <-p.SquereCh:
			if ok {
				fmt.Println("squere channel data:", data)
			} else {
				p.SquereCh = nil
			}
		case data2, ok := <-p.MultiplyeCh:
			if ok {
				fmt.Println("secnd channel data recive:", data2)
			} else {
				p.MultiplyeCh = nil
			}
		}
	}
}
func main() {
	// other
	pool := NewPool()
	pool.wg.Add(3)
	go pool.squere(5)
	go pool.multiply(5)
	// pool.wg.Wait()
	go pool.print()
	// pool.wg.Add(1)
	pool.wg.Wait()
}
