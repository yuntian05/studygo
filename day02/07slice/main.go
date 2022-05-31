package main

import "fmt"

func main() {
	var s1 []int
	var s2 []string

	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	s1 = []int{1, 2, 3, 4}
	s2 = []string{"沙河", "张江"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	fmt.Printf("len:%d cap:%d\n", len(s1), cap(s1))
	fmt.Printf("len:%d cap:%d\n", len(s2), cap(s2))

	a1 := [...]int{1, 2, 3, 4, 5, 6, 7}
	s3 := a1[0:4]
	fmt.Println(s3)

	s4 := a1[:3]
	s5 := a1[3:]
	fmt.Println(s4, s5)
	fmt.Printf("len(s4):%d cap(s4):%d\n", len(s4), cap(s4))
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5))
}
