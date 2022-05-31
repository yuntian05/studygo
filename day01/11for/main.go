package main

import "fmt"

func main() {

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// pkg i = 5
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	str := "hello 朝阳"
	for k, v := range str {
		fmt.Printf("%d %c\n", k, v)
	}
}
