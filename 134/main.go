package main

import "fmt"

// 方法1：暴力算法，每个点都尝试下
func canCompleteCircuit(gas []int, cost []int) int {
	for index, value := range gas {
		j := index
		remain := value
		for remain-cost[j] >= 0 {
			remain = remain - cost[j] + gas[(j+1)%len(gas)]
			j = (j + 1) % len(gas)
			if j == index {
				return index
			}
		}
	}
	return -1
}

// 方法2：只要总油量大于消耗油量，就一定存在解，可以证明的；
// 如果选择站点 i 作为起点「恰好」无法走到站点 j，那么 i 和 j 中间的任意站点 k 都不可能作为起点。
// 画图就可以知道了
func canCompleteCircuit1(gas []int, cost []int) int {
	sum := 0
	for index := range gas {
		sum += gas[index] - cost[index]
	}
	if sum < 0 {
		return -1 // 总油量小于消耗油量
	}

	tank, start := 0, 0
	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			// 无法从 start 到达 i + 1
			// 所以站点 i + 1 应该是起点
			start = i + 1
			tank = 0
		}
	}
	return start
}

func main() {
	fmt.Println(canCompleteCircuit1([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
}
