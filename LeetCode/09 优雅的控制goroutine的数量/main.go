package main

import (
	"runtime"
	"sync"
	"time"
)

// Pool 使用channel和sync.WaitGroup实现goroutine池
type Pool struct {
	channel chan int
	sg      *sync.WaitGroup
}

func New(size int) *Pool {
	return &Pool{
		channel: make(chan int, size),
		sg:      &sync.WaitGroup{},
	}
}

func (p *Pool) Add(num int) {
	if num <= 0 {
		panic("num must > 0")
	}
	for i := 0; i < num; i++ {
		p.channel <- 1
	}
	p.sg.Add(num)
}

func (p *Pool) Done() {
	p.sg.Done()
	<-p.channel
}

func (p *Pool) Wait() {
	p.sg.Wait()
}

func main() {
	pool := New(100)
	println(runtime.NumGoroutine())
	for i := 0; i < 1000; i++ {
		pool.Add(1)
		go func() {
			time.Sleep(time.Second)
			println(runtime.NumGoroutine())
			pool.Done()
		}()
	}
	pool.Wait()
	println(runtime.NumGoroutine())
}
