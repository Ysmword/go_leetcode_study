package main

import (
	"fmt"
	"math"
)

// 哈希表搞定
func containsNearbyDuplicate(nums []int, k int) bool {
	nMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		index, ok := nMap[nums[i]]
		if math.Abs(float64(i)-float64(index)) <= float64(k) && i != 0 && ok {
			return true
		}
		nMap[nums[i]] = i
	}
	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
