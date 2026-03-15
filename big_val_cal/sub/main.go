// 有符号相减
package main

import "fmt"

func dispatch(num1, num2 string) string {

	f1, f2 := string(num1[0]), string(num2[0])
	if f1 == "-" && f2 != "-" {
		return "-" + bigSum(num1[1:], num2)
	}
	if f1 == "-" && f2 == "-" {
		return bigSub(num2[1:], num1[1:])
	}
	if f1 != "-" && f2 == "-" {
		return bigSum(num1, num2[1:])
	}
	if f1 != "-" && f2 != "-" {
		return bigSub(num1, num2)
	}
	return ""
}

func bigSum(num1, num2 string) string {
	carry := 0
	x, y := len(num1)-1, len(num2)-1
	sum := ""
	for x >= 0 && y >= 0 {
		n1 := int(num1[x] - '0')
		n2 := int(num2[y] - '0')
		val := n1 + n2 + carry
		v := val % 10
		carry = val / 10
		sum = fmt.Sprintf("%d", v) + sum
		x--
		y--
	}
	if x >= 0 {
		n1 := int(num1[x] - '0')
		val := n1 + carry
		v := val % 10
		carry = val / 10
		sum = fmt.Sprintf("%d", v) + sum
		x--
	}
	if y >= 0 {
		n2 := int(num2[y] - '0')
		val := n2 + carry
		v := val % 10
		carry = val / 10
		sum = fmt.Sprintf("%d", v) + sum
		y--
	}
	return sum
}

func bigSub(num1, num2 string) string {
	borrow := 0
	x, y := len(num1)-1, len(num2)-1
	sub := ""
	for x >= 0 && y >= 0 {
		n1 := int(num1[x] - '0')
		n2 := int(num2[y] - '0')
		val := n1 - n2 - borrow
		if val < 0 {
			val += 10
			borrow = 1
		} else {
			borrow = 0
		}
		sub = fmt.Sprintf("%d", val) + sub
		x--
		y--
	}

	if x >= 0 {
		n1 := int(num1[x] - '0')
		val := n1 - borrow
		if val < 0 {
			val += 10
			borrow = 1
		} else {
			borrow = 0
		}
		sub = fmt.Sprintf("%d", val) + sub
		x--
	}

	if y >= 0 {
		n2 := int(num2[y] - '0')
		val := n2 - borrow
		if val < 0 {
			val += 10
			borrow = 1
		} else {
			borrow = 0
		}
		sub = fmt.Sprintf("%d", val) + sub
		y--
	}
	// 去掉前导0
	newSub := ""
	for i := 0; i < len(sub); i++ {
		if sub[i] == '0' {
			continue
		}
		newSub = newSub + sub
	}
	if newSub == "" {
		return "0"
	}

	if borrow > 0 {
		newSub = "-" + newSub
	}

	return newSub
}

func main() {
	fmt.Println(dispatch("-11", "11"))
	fmt.Println(dispatch("-11", "-11"))
	fmt.Println(dispatch("11", "-11"))
	fmt.Println(dispatch("11", "11"))

}
