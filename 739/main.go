package main

import (
	"fmt"
	"reflect"
)

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0) // 单调递减栈，存放的是温度的下标
	answer := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			answer[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	// 剩余的温度，都没有找到之后更大的,这个部分是不需要的，因为没有赋值就是为0
	for len(stack) != 0 {
		answer[stack[len(stack)-1]] = 0
		stack = stack[:len(stack)-1]
	}
	return answer
}

func main() {
	ans := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	fmt.Println(ans)

	if reflect.DeepEqual(ans, []int{1, 1, 4, 2, 1, 1, 0, 0}) {
		fmt.Println("pass")
	}

	fmt.Println("=====1========")
	ans = dailyTemperatures1([]int{73, 74, 75, 71, 69, 72, 76, 73})
	fmt.Println(ans)

	if reflect.DeepEqual(ans, []int{1, 1, 4, 2, 1, 1, 0, 0}) {
		fmt.Println("pass")
	}

}

func dailyTemperatures1(temperatures []int) []int {
	res := make([]int, len(temperatures))
	st := make([]int, 0)
	
	for i := 0; i < len(temperatures); i++ {
		for len(st) > 0 && temperatures[st[len(st)-1]] < temperatures[i] {
			index := st[len(st)-1]
			st = st[:len(st)-1]
			res[index] = i - index
		}
		st = append(st, i)
	}
	return res
}
