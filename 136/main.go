package main

import "fmt"

// 异或
func singleNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res = res ^ nums[i]
	}
	return res
}

func main() {
	fmt.Println(singleNumber([]int{1}))
}
