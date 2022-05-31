package main

func quicksort(nums []int, low, high int) {
	if low >= high {
		return
	}

	left := low
	right := high
	pivot := nums[left]
	// 结束条件是left==right
	for left < right {
		// 右边 大于基准值
		for left < right && nums[right] > pivot {
			right--
		}
		if left < right {
			nums[left] = nums[right]
			left++
		}
		// 左边 小于基准值
		for left < right && nums[left] < pivot {
			left++
		}
		if left < right {
			nums[right] = nums[left]
			right--
		}
	}
	// 基准值放到该位置
	nums[left] = pivot

	quicksort(nums, low, left-1)
	quicksort(nums, left+1, high)
}

func findKthLargest(nums []int, k int) int {
	revquicksort(nums, 0, len(nums) - 1)
	for index, val := range nums {
		if index + 1 == k {
			return val
		}
	}
	return -1
}

// 移除相同元素
func movedupElement(nums []int)[]int {
	length := len(nums)
	if length == 0 {
		return nums
	}
	slow,fast := 0, 0
	for fast < length {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return nums[:slow+1]
}

// 逆序
func revquicksort(nums []int, low, high int) {
	if low >= high {
		return
	}

	pivotIndex := partion(nums, low, high)
	revquicksort(nums, low, pivotIndex - 1)
	revquicksort(nums, pivotIndex + 1, high)
}

func partion(nums []int, left, right int) int {
	pivot := nums[left]
	// left == right 结束
	for left < right {
		// 右边 小于于基准值
		for left < right && pivot > nums[right] {
			right--
		}
		if left < right {
			// 大于基准值
			nums[left] = nums[right]
			left++
		}

		// 左边 大于基准值
		for left < right && pivot < nums[left] {
			left++
		}
		if left < right {
			nums[right] = nums[left]
			right--
		}
	}
	nums[left] = pivot
	return left
}
