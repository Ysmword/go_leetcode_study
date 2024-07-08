package main

import "fmt"

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

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
