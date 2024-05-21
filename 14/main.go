package main

import "fmt"

// 方法1: 以任意一个元素作为参考
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	res := ""
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i > len(strs[j])-1 || strs[j][i] != strs[0][i] {
				return res
			}
		}
		res = res + string(strs[0][i])
	}
	return res
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flower", "flower", "flower"}))
}
