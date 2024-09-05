package main

import "fmt"

// 从后面开始遍历，然后如果是9就变成0，然后如果不是9就+1
func plusOne(digits []int) []int {
	carry := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			value := digits[i] + 1
			digits[i] = value % 10
			carry = value / 10
			continue
		}
		value := digits[i] + carry
		digits[i] = value % 10
		carry = value / 10
	}
	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{9, 9, 9}))
	fmt.Println(plusOne([]int{0}))
}
