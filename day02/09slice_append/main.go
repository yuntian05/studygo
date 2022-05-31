package main

import "fmt"

func main() {
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1 = %v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "广州")
	fmt.Printf("s1 = %v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "杭州", "成都")
	fmt.Printf("s1 = %v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	ss := []string{"武汉", "西安", "苏州", "长沙", "无锡", "大连", "金华"}
	s1 = append(s1, ss...)
	fmt.Printf("s1 = %v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
}
