package main

import (
	"fmt"
	"strings"
)

// 方法1：直接调用提供好的方法
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// 方法2：遍历加截取判断是否相等
func strStr1(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			if i+len(needle)-1 <= len(haystack)-1 {
				if haystack[i:i+len(needle)] == needle {
					return i
				}
			} else {
				return -1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr1("sadbutsad", "sad"))
}
