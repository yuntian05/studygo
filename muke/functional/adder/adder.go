package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(val int) (int, iAdder) {
		return base + val, adder2(base + val)
	}
}

func main() {
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		//fmt.Printf("0 + 1 + 2 + ... + %d = %d\n", i, a(i))
		fmt.Printf("0 + 1 + 2 + ... + %d = %d\n", i, s)
	}
}
