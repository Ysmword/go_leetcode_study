package main

import (
	"fmt"
)

// 不懂
func reverseBits(num uint32) uint32 {
	res := uint32(0)
	for i := 0; i < 32 && num > 0; i++ {
		// (num & 1) 每次取最后一位
		res |= (num & 1) << (31 - i)
		num >>= 1
	}
	return res
}

// 思路很简单：找到第i位，移动到31-i位，i从0开始
func reverseBits1(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		bit := (num >> i) & 1 // 第i位的值
		res = res | (bit << (31 - i))
	}
	return res
}

func main() {
	fmt.Println(reverseBits1(10))

	// fmt.Printf("%b\n", 1)
	// fmt.Printf("%b\n", 1<<1)
	// fmt.Printf("%b\n", 1<<2)
	// fmt.Printf("%b\n", 1<<8)

	// fmt.Printf("%b\n", 1<<8)
}
