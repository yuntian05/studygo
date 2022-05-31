package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for n := range c{
		fmt.Printf("worker %d received %d\n", id, n)
	}
}
func creatWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}
func channelDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = creatWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Second)
}

func channelClose() {
	c := make(chan int, 3)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	go worker(0, c)
	go worker(0, c)
	go worker(0, c)
	close(c)
	time.Sleep(time.Second)
}
func main() {
	//channelDemo()
	channelClose()
}
