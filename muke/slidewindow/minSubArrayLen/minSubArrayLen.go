package main

func minSubArrayLen(nums []int, target int) int {
	len1 := len(nums)
	if len1 == 0 {
		return 0
	}

	windoSum := 0
	left, right := 0, 0
	ret := len1 + 1
	change := false
	for right < len1 {
		// 移入窗口
		windoSum += nums[right]
		right++

		// 缩小窗口
		for windoSum >= target && left < right{
			if right - left < ret {
				ret = right - left
				change = true
			}
			// 移出窗口
			windoSum -= nums[left]
			left++
		}
	}
	if change {
		return ret
	} else {
		return 0
	}
}
