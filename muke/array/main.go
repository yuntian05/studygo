package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	//fmt.Println(removeDuplicates(nums))
	//fmt.Println(removeElement(nums, 2))
	//moveZeroes(nums)
	fmt.Println(binarySearch(nums, 7))
	fmt.Println(binarySearchLeftBound(nums, 6))
	fmt.Println(binarySearchRightBound(nums, 6))
}

// 删除重复元素
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			// 维护 nums[0..slow] 无重复
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

// 移除指定元素
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func moveZeroes(nums []int) {
	p := removeElement(nums, 0)
	for ; p < len(nums); p++ {
		nums[p] = 0
	}
}

func binarySearch(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			return mid
		}
	}
	return -1
}

// 二分查找左边界
func binarySearchLeftBound(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 收缩右边界
			right = mid - 1
		}
	}

	if left >= length || nums[left] != target {
		return -1
	}
	return left
}

func binarySearchRightBound(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 收缩左边界
			left = mid + 1
		}
	}

	if right < 0 || nums[right] != target {
		return -1
	}

	return right
}
