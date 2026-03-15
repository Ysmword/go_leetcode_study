package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	// 先转化为字符串
	numStrs := make([]string, len(nums))
	for i, num := range nums {
		numStrs[i] = strconv.Itoa(num)
	}

	// 排序
	sort.SliceStable(numStrs, func(i, j int) bool {
		// 字符串判断，第一位越大的，字符串越大
		return numStrs[i]+numStrs[j] > numStrs[j]+numStrs[i]
	})

	// 处理前导0
	if numStrs[0] == "0" {
		return "0"
	}

	return strings.Join(numStrs, "")
}

func main() {
	// nums := []int{10, 2}
	// fmt.Println(largestNumber(nums))

	nums := []int{3, 30, 34, 5, 9}
	fmt.Println(largestNumber(nums))

	fmt.Println("====1======")
	fmt.Println(largestNumber1(nums))
}

func largestNumber1(nums []int) string {
	numStr := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		numStr[i] = fmt.Sprintf("%d", nums[i])
	}
	sort.SliceStable(numStr, func(i, j int) bool {
		return numStr[i]+numStr[j] > numStr[j]+numStr[i]
	})
	return strings.Join(numStr, "")
}
