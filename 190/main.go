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

func main() {
	fmt.Println(reverseBits(10))

	// fmt.Printf("%b\n", 1)
	// fmt.Printf("%b\n", 1<<1)
	// fmt.Printf("%b\n", 1<<2)
	// fmt.Printf("%b\n", 1<<8)

	// fmt.Printf("%b\n", 1<<8)
}
