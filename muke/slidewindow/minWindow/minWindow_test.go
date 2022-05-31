package main

import (
	"testing"
)

func TestMiniWindow(t *testing.T) {
	tests := []struct{str, target , ret string} {
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"", "ABC", ""},
		{"", "", ""},
		{"a", "a", "a"},
	}

	for _, tt := range tests {
		actual := minWindow(tt.str, tt.target)
		if actual != tt.ret {
			t.Errorf("got:%s, expected:%s\n", actual, tt.ret)
		}
	}
}
