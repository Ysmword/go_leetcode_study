package main

import (
	"fmt"
)

// 不懂
func reverseBits(num uint8) uint8 {
	var res uint8
	for i := 0; i < 8 && num > 0; i++ {
		// fmt.Printf("tmp1:%b\n", 1<<(8-i))
		// fmt.Printf("tmp2:%b\n", num)
		// fmt.Printf("tmp3:%b\n", num&1<<(8-i))
		res |= num & 1 << (8 - i) // 先移动n位置，看num的n位置是否为1或者0
		num >>= 1                 // num移动1位
	}
	fmt.Printf("%b\n", res)
	return res
}

// 0000 0000 0000 0000 0000 0000 0000 1010  表示为10
// res |= num & 1 << (31 - i) 计算，
/*
当i为0时，1<<31 为1000 0000 0000 0000 0000 0000 0000 0000
num & 1 << 31  为0000 0000 0000 0000 0000 0000 0000 0000
res |= num & 1 << 31 为0000 0000 0000 0000 0000 0000 0000 0000

当i为1时，1<<30为      0100 0000 0000 0000 0000 0000 0000 0000
num 为                0000 0000 0000 0000 0000 0000 0000 0101
num & 1 << 30       为0000 0000 0000 0000 0000 0000 0000 0000
res |= num & 1 << 30为0000 0000 0000 0000 0000 0000 0000 0000
*/

func main() {
	fmt.Println(reverseBits(10))

	// fmt.Printf("%b\n", 1)
	// fmt.Printf("%b\n", 1<<1)
	// fmt.Printf("%b\n", 1<<2)
	// fmt.Printf("%b\n", 1<<8)

	// fmt.Printf("%b\n", 1<<8)
}
