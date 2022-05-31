package main

import "testing"

func TestQuicksort(t *testing.T) {
	tests := []struct {
		nums []int
		ans  []int
	}{
		{[]int{5, 9, 1, 9, 5, 3, 7, 6, 1}, []int{1, 1, 3, 5, 5, 6, 7, 9, 9}},
	}

	for _, tt := range tests {
		if len(tt.ans) != len(tt.nums) {
			t.Errorf("数组大小不等")
			continue
		}
		quicksort(tt.nums, 0, len(tt.nums)-1)
		for index, val := range tt.nums {
			if tt.ans[index] != val {
				t.Errorf("got:%v, expected:%v\n", tt.nums, tt.ans)
				break
			}
		}
	}

	tests1 := []struct {
		nums []int
		k int
		ans  int
	}{
		{[]int{3,2,3,1,2,4,5,5,6}, 4, 4},
	}
	for _, tt := range tests1 {
		actual := findKthLargest(tt.nums, tt.k)
		if actual != tt.ans {
			t.Errorf("sort arr:%v got:%d, expected:%d\n", tt.nums, actual, tt.ans)
		}
	}
}
