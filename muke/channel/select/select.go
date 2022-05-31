package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("worker %d received %d\n", id, n)
	}
}
func creatWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Microsecond)
			out <- i
			i++
		}
	}()
	return out
}
func main() {
	var c1, c2 = generator(), generator()
	w := creatWorker(0)
	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			//fmt.Println("received from c1:", n)
			//w <- n
			values = append(values, n)
		case n := <-c2:
			//fmt.Println("received from c2:", n)
			//w <- n
			//hasValue = true
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len:", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
