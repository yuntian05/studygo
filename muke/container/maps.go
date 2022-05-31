package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(courseName)

	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key cause does not exist")
	}

	fmt.Println("deleting values")
	if name, ok := m["name"]; ok {
		fmt.Println(name)
		delete(m, "name")
	} else {
		fmt.Println("key name does not exist")
	}

	if name, ok := m["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("key name does not exist")
	}

	//nums := []int {0, 0, 1, 1, 2, 2, 3, 4}
	//removeDuplicates(nums)
	//removeElement(nums, 2)
	//moveZeros(nums)
	//s := "ADOBECODEBANC"
	//s := "a"
	////t := "ABC"
	//t := "a"
	//minWindow(s, t)
	nums := []int{2, 7, 11, 15}
	twoSum(nums, 9)
}

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	slow := 0
	fast := 0
	for fast < length {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	newArrLen := slow + 1
	nums = nums[0:newArrLen]
	return newArrLen
}


func removeElement(nums []int, val int) int{
	length := len(nums)
	if length == 0 {
		return 0
	}

	slow := 0
	fast := 0
	for fast < length {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	nums = nums[0:slow]
	return slow
}

func moveZeros(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}

	slow := 0
	fast := 0
	for fast < length {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	for ;slow < length; slow++ {
		nums[slow] = 0
	}
}

func minWindow(s string, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)
	tmpT := []byte(t)
	for _, val :=range tmpT {
		need[val]++
	}

	length := len(s)
	left := 0
	right := 0
	valid := 0
	start := 0
	sLen := length+1
	change := false

	tmpS := []byte(s)
	for right < length {
		// c 是将移入窗口的字符
		c := tmpS[right]
		// 扩大窗口
		right++

		// 进行窗口内数据的一系列更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// 判断左侧窗口是否要收缩
		for valid == len(need) {
			// 在这里更新最小覆盖子串
			if right - left < sLen {
				start = left
				sLen = right - left
				change = true
			}

			// d 是将移出窗口的字符
			d := tmpS[left]
			left++

			// 进行窗口内数据的一系列更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if !change {
		return ""
	}
	end := start + sLen
	tmpS = tmpS[start:end]
	ret := string(tmpS)
	return ret
}

// 两数之和
func twoSum(nums []int, target int)[]int {
	sort.Ints(nums)
	length := len(nums)
	if length == 0 {
		return []int{}
	}
	left := 0
	right := len(nums) - 1
	for left < right {
		numL := nums[left]
		numR := nums[right]
		sum := numL + numR
		if sum == target {
			return []int{left, right}
		} else if sum > target {
			// 让sum小一点
			for left < right && numR == nums[right] {
				right--
			}

		} else {
			// 让sum大一点
			for left < right && numL == nums[left] {
				left++
			}
		}
	}
	return []int{}
}