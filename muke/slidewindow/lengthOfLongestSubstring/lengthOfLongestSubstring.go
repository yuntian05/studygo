package main


func lengthOfLongestSubstring(s string) int {
	r := []rune(s)
	length := len(r)
	if length == 0 {
		return 0
	}
	window := make(map[rune]int)
	left, right := 0,0
	res := 0
	for right < length {
		// c是移入窗口的字符
		c := r[right]
		right++
		window[c]++
		// 收缩窗口
		for window[c] > 1 {
			// d是移出窗口的字符
			d := r[left]
			left++
			window[d]--
		}
		// 更新宽度
		width := right - left
		if res < width {
			res = width
		}
	}
	return res
}
