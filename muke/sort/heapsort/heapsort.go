package main

func heapsort(nums []int, n int) {
	buildheap(nums, n)

	for i := n - 1; i >= 0; i-- {
		// 最后一个数和第一个数互换
		swap(nums, i, 0)
		//
		heapify(nums, i, 0)
	}
}

// 建立大顶堆
func buildheap(nums []int, n int) {
	lastnode := n - 1
	parent := (lastnode - 1) / 2
	for i := parent; i >= 0; i-- {
		heapify(nums, n, i)
	}
}

// 将i下沉
func heapify(nums []int, n, i int) {
	if i >= n {
		return
	}
	left := 2 * i + 1
	right := 2 * i + 2
	max := i
	if left < n && nums[left] > nums[max] {
		max = left
	}
	if right < n && nums[right] > nums[max] {
		max = right
	}
	if max != i {
		swap(nums, max, i)
		heapify(nums, n, max)
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
