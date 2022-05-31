package main

import "testing"

func TestCheckInclusion(t *testing.T) {
	tests := []struct{
		s1 string
		s2 string
		ans bool
	} {
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
		{"aa", "eidboaao", true},
		{"abcdxabcde","abcdeabcdx", true},
	}

	for _, tt := range tests {
		actual := checkInclusion(tt.s1, tt.s2)
		if actual != tt.ans {
			t.Errorf("s1=%s, s2= %s, got:%v, expected:%v", tt.s1, tt.s2, actual, tt.ans)
		}
	}
}
