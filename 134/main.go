package main

// 方法1：gas和cost和相等，就一定能够绕一周
func canCompleteCircuit(gas []int, cost []int) int {
	gSum, cSum := 0, 0
	start := 0
	for i := 0; i < len(gas); i++ {
		gSum += gas[i]
		cSum += cost[i]
		if gas[i] > cost[i] {
			if start == 0 {
				start = i
			}
		}
	}
	if gSum == cSum {
		return start
	}
	return -1
}
