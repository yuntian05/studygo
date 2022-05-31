package main

import "testing"

func TestMergesort(t *testing.T) {
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
		temp := make([]int, len(tt.nums))
		mergesort(tt.nums, temp,0, len(tt.nums)-1)
		for index, val := range tt.nums {
			if tt.ans[index] != val {
				t.Errorf("got:%v, expected:%v\n", tt.nums, tt.ans)
				break
			}
		}
	}
}
