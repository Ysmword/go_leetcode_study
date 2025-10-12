package main

import (
	"fmt"
	"strings"
)

// 判断n1是否大于n2
func compare(n1, n2 string) bool {
	if len(n1) > len(n2) {
		return true
	}
	if len(n1) < len(n2) {
		return false
	}

	for i := 0; i < len(n1); i++ {
		if n1[i] > n2[i] {
			return true
		}
	}
	return false
}

func subtractUnsignedBigNumbers(n1, n2 string) string {
	f := compare(n1, n2)
	if !f {
		n1, n2 = n2, n1
	}

	// 反转字符串，从最低位开始计算
	revNum1 := reverse(n1)
	revNum2 := reverse(n2)

	var res []byte
	borrow := 0

	for i := 0; i < len(revNum1); i++ {
		d1 := int(revNum1[i] - '0')
		d2 := 0
		if i < len(revNum2) {
			d2 = int(revNum2[i] - '0')
		}

		temp := d1 - d2 - borrow
		if temp < 0 {
			borrow = 1
			temp += 10
		} else {
			borrow = 0
		}

		res = append(res, byte(temp)+'0')
	}

	// 反转结果，恢复正确顺序
	resStr := reverse(string(res))
	// 去除前导零（若结果全零则保留一个零）
	resStr = strings.TrimLeft(resStr, "0")
	if resStr == "" {
		return "0"
	}
	if !f {
		resStr = "-" + resStr
	}
	return resStr
}

// reverse 反转字符串（用于从最低位开始计算）
func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {
	res := subtractUnsignedBigNumbers("12", "13")
	fmt.Println(res)
}
