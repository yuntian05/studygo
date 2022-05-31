package main

import "testing"

func TestHeapify(t *testing.T) {
	tests := []struct {
		nums []int
		ans  []int
	}{
		{[]int{4, 10, 3, 5, 1, 2}, []int{}},
	}

	for _, tt := range tests {
		heapify(tt.nums, len(tt.nums), 0)
	}
}

func TestBuildheap(t *testing.T) {
	tests := []struct {
		nums []int
		ans  []int
	}{
		{[]int{2, 5, 3, 1, 10, 4}, []int{10, 5, 4, 1, 2, 3}},
		{[]int{2, 5, 3, 1, 10, 4}, []int{10, 5, 4, 1, 2, 3}},
	}

	for _, tt := range tests {
		buildheap(tt.nums, len(tt.nums))
		t.Errorf("got:%v, expected:%v\n", tt.nums, tt.ans)
	}
}
func TestHeapsort(t *testing.T) {
	tests := []struct {
		nums []int
		ans  []int
	}{
		{[]int{2, 5, 3, 1, 10, 4}, []int{}},
	}

	for _, tt := range tests {
		heapsort(tt.nums, len(tt.nums))
	}
}
