package main

import "fmt"

// 方法1：利用额外空间map和减法
func twoSum(numbers []int, target int) []int {
	nM := make(map[int][]int)
	for index, n := range numbers {
		if _, ok := nM[target-n]; ok {
			return []int{nM[target-n][0], index + 1}
		}
		nM[n] = append(nM[n], index+1)
	}
	return nil
}

// 方法2：使用双指针
func twoSum1(numbers []int, target int) []int {
	x, y := 0, len(numbers)-1
	for x < y {
		if numbers[x]+numbers[y] < target {
			x++
		} else if numbers[x]+numbers[y] > target {
			y--
		} else {
			return []int{x + 1, y + 1}
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 3, 4}, 6))

	fmt.Println(twoSum1([]int{2, 3, 4}, 6))
}
