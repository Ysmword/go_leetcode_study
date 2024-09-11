package main

import (
	"fmt"
	"strconv"
)

// 简单思路，但是要注意要倒着排序
func addBinary(a string, b string) string {
	x, y := len(a)-1, len(b)-1
	res := ""
	tempValue := ""
	for x >= 0 && y >= 0 {
		if a[x] == b[y] && a[x] == '1' {
			if tempValue == "" {
				res = "0" + res
				tempValue = "1"
			} else {
				res = "1" + res
				tempValue = "1"
			}
		} else if (a[x] == '0' && b[y] == '1') || (a[x] == '1' && b[y] == '0') {
			if tempValue == "" {
				res = "1" + res
			} else {
				res = "0" + res
				tempValue = "1"
			}
		} else {
			if tempValue == "" {
				res = "0" + res
			} else {
				res = "1" + res
				tempValue = ""
			}
		}
		x--
		y--
	}

	for x >= 0 {
		if tempValue == "" {
			res = a[:x+1] + res
			break
		} else {
			if a[x] == '0' {
				res = "1" + res
				tempValue = ""
			} else {
				res = "0" + res
				tempValue = "1"
			}
		}
		x--
	}

	for y >= 0 {
		if tempValue == "" {
			res = b[:y+1] + res
			break
		} else {
			if b[y] == '0' {
				res = "1" + res
				tempValue = ""
			} else {
				res = "0" + res
				tempValue = "1"
			}
		}
		y--
	}
	if tempValue != "" {
		return tempValue + res
	}

	return res
}

// 简介做法, 巧妙运用carry
func addBinary1(a string, b string) string {
	ans := ""
	carry := 0
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)

	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry = carry / 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

func main() {
	// fmt.Println(addBinary("11", "1"))
	// fmt.Println(addBinary("1010", "1011"))
	// fmt.Println(addBinary("100", "110010"))
	fmt.Println(addBinary("110010", "10111"))
}
