package main

import "fmt"

// 双指针=》滑动窗口，当出现重复的时候，左边指针动
func lengthOfLongestSubstring(s string) int {
	subStrMap := make(map[string]int)
	x, y := 0, 0
	maxSub := 0
	for y < len(s) {
		index, ok := subStrMap[string(s[y])]
		if ok && index >= x { // 存在，并且存在的下标大于等于x的时候，计算不出现重复的子字符串情况
			// 这一步很妙，因为不用delete，index < x 表示虽然出现过，但是不是【x,y】之间的连续字符串
			maxSub = max(maxSub, y-x)
			x = index + 1
		}
		subStrMap[string(s[y])] = y
		y++
	}
	return max(maxSub, y-x)
}

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(" "))
	fmt.Println(lengthOfLongestSubstring("abba"))
}
