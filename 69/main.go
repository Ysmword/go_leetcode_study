package main

import "fmt"

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	low, high := 1, x/2
	for low <= high {
		mid := (low + high) / 2
		val := mid * mid
		if x == val {
			return mid
		} else if x < val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low - 1
}

func main() {
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(9))
	fmt.Println(mySqrt(1))
	fmt.Println(mySqrt(2))
}
