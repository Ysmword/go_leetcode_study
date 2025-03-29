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

/*
首先判断总gas能不能大于等于总cost，如果总gas不够，一切都白搭对吧（总（gas- cost）不用单独去计算，和找最低点时一起计算即可，只遍历一次）；

再就是找总（gas-cost）的最低点，不管正负（当然如果最低点都是正的话那肯定能跑完了）；

找到最低点后，如果有解，那么解就是最低点的下一个点，因为总（gas-cost）是大于等于0的，所以前面损失的gas我从最低点下一个点开始都会拿回来！（此处@小马哥！“我要争一口气，不是想证明我了不起。我是要告诉人家，我失去的东西一定要拿回来！”），别管后面的趋势是先加后减还是先减后加，最终结果我是能填平前面的坑的。
*/
func canCompleteCircuit2(gas []int, cost []int) int {
	curSum, totalSum, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		if curSum < 0 {
			start = i + 1
			curSum = 0
		}
	}
	if totalSum < 0 {
		return -1
	}
	return start
}

func main() {
	fmt.Println(canCompleteCircuit1([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
}
