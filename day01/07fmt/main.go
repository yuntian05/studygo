package main

import "fmt"

func main() {
	num := 100
	fmt.Printf("%d\n", num)
	fmt.Printf("%T\n", num)

	str := "hello 朝阳"
	fmt.Printf("%s\n", str)
	fmt.Printf("%T\n", str)
	fmt.Printf("%#v\n", str)
}
