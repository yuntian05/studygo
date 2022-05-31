package main

import (
	"fmt"
	"sync"
)
type worker struct {
	in chan int
	done func()
}

func doworke(id int, w worker) {
	for n := range w.in{
		fmt.Printf("doworke %d received %d\n", id, n)
		w.done()
	}
}
func creatWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doworke(id, w)
	return w
}
func channelDemo() {
	var channels [10]worker
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		channels[i] = creatWorker(i, wg)
	}
	wg.Add(20)
	for i, w :=range channels{
		w.in <- 'a' + i
	}

	for i, w :=range channels{
		w.in <- 'A' + i
	}
	wg.Wait()
}

func main() {
	channelDemo()
}
