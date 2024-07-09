package main

import (
	"fmt"
	"strings"
)

// 方法1：切割后，在倒叙遍历后连接
func reverseWords(s string) string {
	items := strings.Split(s, " ")
	result := make([]string, 0)
	for i := len(items) - 1; i >= 0; i-- {
		if items[i] == "" {
			continue
		}
		result = append(result, items[i])
	}
	return strings.Join(result, " ")
}

// 方法2：倒叙遍历


func main() {
	fmt.Println(reverseWords("a good   example"))
}
