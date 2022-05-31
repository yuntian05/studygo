package main

import (
	"testing"
)

func TestSubstring(t *testing.T) {

	tests := []struct {
		str string
		num int
	}{
		// normal case
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// edage case
		{"", 0},
		{"b", 1},
		{"bbbb", 1},
		{"abcabcabcd", 4},

		// chinese case
		{"这里是慕课网", 6},
	}

	for _, tt := range tests {
		actual := lengthOfLongestSubstring(tt.str)
		if actual != tt.num {
			t.Errorf("got:%d, expected:%d\n", actual, tt.num)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	ans := 8

	for i := 0; i < b.N; i++ {
		actual := lengthOfLongestSubstring(s)
		if actual != ans {
			b.Errorf("got:%d, expected:%d\n", actual, ans)
		}
	}
}
