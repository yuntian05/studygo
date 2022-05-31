package main

import "fmt"

func main() {
	// pkg a1 [3]bool
	a1 := [3]bool{true, false, false}
	fmt.Println(a1)

	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a10)

	a5 := [5]int{0: 1, 4: 2}
	fmt.Println(a5)

	citys := []string{"北京", "上海", "深圳"}
	for _, city := range citys {
		fmt.Println(city)
	}

	var a6 [3][2]string
	a6 = [3][2]string{
		[2]string{"北京", "上海"},
		[2]string{"广州", "深圳"},
		[2]string{"成都", "重庆"},
	}
	for _, v1 := range a6 {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
}
