package main

import "fmt"

// 回溯算法
func generateParenthesis(n int) []string {
	var res []string
	backtrace("(", 1, 0, n, &res)
	return res
}

func backtrace(s string, left, right, n int, res *[]string) {
	if left == right && len(s) == 2*n {
		*res = append(*res, s)
		return
	}
	if left > n || right > n {
		return
	}
	if left > right {
		s += ")"
		backtrace(s, left, right+1, n, res)
		s = s[:len(s)-1]
	}
	s += "("
	backtrace(s, left+1, right, n, res)
}

func generateParenthesis1(n int) []string {
	var res []string
	backtrace1("", 0, 0, n, &res)
	return res
}

func backtrace1(s string, left, right, n int, res *[]string) {
	if left == right && left == n {
		*res = append(*res, s)
		return
	}

	if left > n || right > n {
		return
	}

	if left < right { // 右括号太多，组成的肯定无效
		return
	}

	if left > right {
		s += ")"
		backtrace1(s, left, right+1, n, res)
		s = s[:len(s)-1]
	}
	s += "("
	backtrace1(s, left+1, right, n, res) // 这个时候，我没有必要在回溯，因为我就是要选择他
}

/*
画个图就能知道了
*/

func main() {
	fmt.Println(generateParenthesis(3))
}
