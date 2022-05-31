package main

import "fmt"

func main() {
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1)
	fmt.Printf("%o\n", i1)
	fmt.Printf("%x\n", i1)
	i2 := 077
	fmt.Printf("%d\n", i2)
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)

	// 查看变量的类型
	fmt.Printf("%T\n", i3)

	i4 := int8(9)
	fmt.Printf("%T\n", i4)
}
