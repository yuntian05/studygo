package main

import "fmt"

func minWindow(s string, t string) string {
	r := []rune(s)
	length := len(r)
	if length == 0 || len(s) == 0 {
		return ""
	}
	// 初始化need window
	need := make(map[rune]int)
	window := make(map[rune]int)
	for _, val := range t {
		need[val]++
	}

	left, right := 0, 0
	// 记录有效字符数量
	valid := 0
	needNum := len(need)
	start, strLen := 0, length+1
	change := false
	for right < length {
		// c是将移入窗口的字符
		c := r[right]
		right++

		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		fmt.Printf("window:[%d, %d]\n", left, right)

		// 缩小窗口
		for valid == needNum {
			if right-left < strLen {
				// 最小覆盖窗口
				start = left
				strLen = right - left
				change = true
			}
			// d是将移出窗口的字符
			d := r[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if change {
		return string(r[start : start+strLen])
	} else {
		return ""
	}
}
