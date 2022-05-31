package main

import (
	"fmt"
	"sort"
)

type OrderByLengthDesc []string

func (s OrderByLengthDesc) Len() int {
	return len(s)
}

func (s OrderByLengthDesc) Swap(i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}

func (s OrderByLengthDesc) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func main() {
	city := []string{"New York", "London", "Washington", "Delhi"}
	sort.Sort(OrderByLengthDesc(city))
	fmt.Println(city)

	s := []byte("hello")
	reverseString(s)
	fmt.Println(string(s))

	fmt.Println(isPalindrome("aba"))
	fmt.Println(isPalindrome("abbs"))

	fmt.Println(longestPalindrome("babad"))
}

func reverseString(s []byte) {
	length := len(s)
	if length == 0 {
		return
	}
	left, right := 0, length-1
	for left < right {
		tmp := s[left]
		s[left] = s[right]
		s[right] = tmp
		left++
		right--
	}

}

// 判断是否是回文串
func isPalindrome(s string) bool {
	length := len(s)
	if length == 0 {
		return false
	}

	left, right := 0, length-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 返回回文串 (以left right为中心 向两边扩散)
func palindrome(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

// 最长回文子串
func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		// 以 s[i] 为中心的最长回文子串
		s1 := palindrome(s, i, i)
		// 以 s[i] 和 s[i+1] 为中心的最长回文子串
		s2 := palindrome(s, i, i+1)
		if len(res) < len(s1) {
			res = s1
		}
		if len(res) < len(s2) {
			res = s2
		}
	}
	return res
}
