package main

import "fmt"

var data = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

// 暴力算法
func intToRoman(num int) string {
	result := ""
	if num >= 1000 {
		for i := 1; i <= num/1000; i++ {
			result += "M"
		}
		num = num % 1000
	}
	if 100 <= num && num < 1000 {
		if num/100 == 4 {
			result += "CD"
		} else if num/100 == 9 {
			result += "CM"
		} else {
			if num >= 500 {
				result += "D"
				num = num - 500
			}
			if num/100 == 4 {
				result += "CD"
			} else {
				for i := 1; i <= num/100; i++ {
					result += "C"
				}
			}
		}
		num = num % 100
	}
	if 10 <= num && num <= 99 {
		if num/10 == 4 {
			result += "XL"
		} else if num/10 == 9 {
			result += "XC"
		} else {
			if num >= 50 {
				result += "L"
				num = num - 50
			}
			if num/10 == 4 {
				result += "XL"
			} else {
				for i := 1; i <= num/10; i++ {
					result += "X"
				}
			}
		}
		num = num % 10
	}
	if num == 9 {
		result += "IX"
	} else if num == 4 {
		result += "IV"
	} else {
		if num >= 5 {
			num -= 5
			result += "V"
		}
		if num == 4 {
			result += "IV"
		} else {
			for i := 1; i <= num; i++ {
				result += "I"
			}
		}
	}
	return result
}

func main() {
	fmt.Println(intToRoman(99))
}
