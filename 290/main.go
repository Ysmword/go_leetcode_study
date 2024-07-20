package main

import (
	"fmt"
	"strings"
)

// 分割+两个哈希表+相互看【看当前值】
func wordPattern(pattern string, s string) bool {
	items := strings.Split(s, " ")
	pMap := make(map[byte]string)
	sMap := make(map[string]byte)
	if len(items) != len(pattern) {
		return false
	}

	for i := range pattern {
		item := pMap[pattern[i]]
		if item != "" && item != items[i] {
			return false
		}
		p := sMap[items[i]]
		if p > 0 && p != pattern[i] {
			return false
		}
		pMap[pattern[i]] = items[i]
		sMap[items[i]] = pattern[i]
	}
	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}
