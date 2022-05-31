package main

import "fmt"

func checkInclusion(s1 string, s2 string)bool {
	r1 := []rune(s1)
	r2 := []rune(s2)
	len1 := len(r1)
	len2 := len(r2)
	if len1 == 0 || len2 == 0 {
		return false
	}

	need := make(map[rune]int)
	window := make(map[rune]int)
	for _, val := range r1 {
		need[val]++
	}
	len3 := len(need)

	left, right := 0, 0
	valid := 0
	for right < len2 {
		// c是将移入窗口的字符
		c := r2[right]
		right++

		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		fmt.Printf("window:[%d, %d]\n", left, right)

		// 收缩窗口
		for right - left >= len1 {
			if valid == len3 {
				return true
			}
			// d是将移出窗口的字符
			d := r2[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	return false
}
