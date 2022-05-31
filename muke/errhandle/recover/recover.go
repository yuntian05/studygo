package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("error occured:", err)
		} else {
			panic(fmt.Sprintf("I do not know how to do:%v", r))
		}
	}()
	//panic(errors.New("this is an error"))
	a := 0
	b := 5/a
	fmt.Println(b)
}
func main() {
	tryRecover()
}
