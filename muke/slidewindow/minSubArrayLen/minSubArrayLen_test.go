package main

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		ans    int
	}{
		//{[]int{2, 3, 1, 2, 4, 3}, 7, 2},
		//{[]int{1, 4, 4}, 4, 1},
		//{[]int{1, 1, 1, 1, 1, 1, 1, 1}, 11, 0},
		{[]int{1, 2, 3, 4, 5}, 11, 3},
	}

	for _, tt := range tests {
		actual := minSubArrayLen(tt.nums, tt.target)
		if actual != tt.ans {
			t.Errorf("nums = %v, target = %d, got %d, expected %d ", tt.nums, tt.target, actual, tt.ans)
		}
	}
}
