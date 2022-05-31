package main


func mergesort(nums []int, temp []int, lo, hi int) {
	if lo >= hi {
		return
	}

	mid := lo + (hi-lo)/2
	mergesort(nums, temp, lo, mid)
	mergesort(nums,temp, mid+1, hi)
	merge(nums, temp, lo, mid, hi)
}

func merge(nums []int, temp []int, lo, mid, hi int) {

	for i := lo; i <= hi; i++ {
		temp[i] = nums[i]
	}

	i, j := lo, mid+1
	for p := lo; p <= hi; p++ {
		if i == mid + 1 {
			// 左半边数组已经合并完毕
			nums[p] = temp[j]
			j++
		} else if j == hi + 1 {
			// 右半边数组已经合并完毕
			nums[p] = temp[i]
			i++
		} else if temp[i] > temp[j] {
			nums[p] = temp[j]
			j++
		} else {
			nums[p] = temp[i]
			i++
		}
	}
}