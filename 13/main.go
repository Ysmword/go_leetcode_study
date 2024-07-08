package main

import (
	"fmt"
)

var data = map[string]int{
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
	"I":  1,
	"V":  5,
	"X":  10,
	"L":  50,
	"C":  100,
	"D":  500,
	"M":  1000,
}

// 方法1：一次遍历，先看两个字符串，在看单个字符串
func romanToInt(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if val, ok := data[s[i:i+2]]; ok {
				result += val
				i += 1
				continue
			}
		}
		result += data[s[i:i+1]]
	}
	return result
}

var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

// 方法2：先看两个，如果左边小于右边，则减少；左边大于等于右边，则增加
func romanToInt2(s string) (ans int) {
	n := len(s)
	for i := range s {
		value := symbolValues[s[i]]
		if i < n-1 && value < symbolValues[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return
}

/*
方法1和方法2只有在输入是正确罗马数字的情况下输出才是正确的结果
*/

func main() {
	fmt.Println(romanToInt("CL"))
	fmt.Println(romanToInt2("CL"))
}
