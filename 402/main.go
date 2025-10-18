package main

import "fmt"

// 贪心 + 单调递增栈 寻找每一位的最小数字
func removeKdigits(num string, k int) string {
	stack := make([]byte, 0)
	for i := 0; i < len(num); i++ {

		// k > 0 表示还可以删除数字；像1432219 为什么还能求出1219, 是因为到1的时候，原本要出两个2，但是k=0了，就出不了了
		for len(stack) > 0 && k > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}

	fmt.Println(string(stack))
	if k > 0 {
		stack = stack[:len(stack)-k]
	}
	// 移除前导0
	for len(stack) > 0 && stack[0] == '0' {
		stack = stack[1:]
	}
	if len(stack) == 0 {
		return "0"
	}
	return string(stack)
}

func main() {
	num := "1432219"
	k := 3
	res := removeKdigits(num, k)
	fmt.Println(res)

	// res = removeKdigits("9", 1)
	// fmt.Println(res)

	// res = removeKdigits("112", 1)
	// fmt.Println(res)
}
