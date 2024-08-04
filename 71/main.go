package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// 调用库函数
func simplifyPath(path string) string {
	res, _ := filepath.Abs(path)
	return res
}

// 分割
func simplifyPath1(path string) string {
	stack := make([]string, 0)
	paths := strings.Split(path, "/")
	for i := 0; i < len(paths); i++ {
		switch paths[i] {
		case "..":
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		case "", ".":
			// 当前目录不做任何处理
		default:
			stack = append(stack, paths[i])
		}
	}
	return "/" + strings.Join(stack, "/")
}

func main() {
	fmt.Println(simplifyPath("/a/./b/../../c/"))

	fmt.Println(simplifyPath1("/a/./b/../../c/"))
}
