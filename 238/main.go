package main

// 方法1: 除法，需要注意几个特殊点，但0个数大于1时，结果都为0，如果0个数等于1，只有0处的值不为0
func productExceptSelf(nums []int) []int {
	sum, count := 1, 0
	for _, num := range nums {
		if num == 0 {
			count++
		} else {
			sum = sum * num
		}
	}
	if count > 1 { // 如果0个数大于1，结果都会是0
		return make([]int, len(nums))
	}
	result := make([]int, len(nums))
	if count == 1 {
		for index, num := range nums {
			if num == 0 {
				result[index] = sum
			}
		}
		return result
	}

	for index, num := range nums {
		if num != 0 {
			result[index] = sum / num
		} else {
			result[index] = sum
		}
	}
	return result
}

// 方法2：可以看作左边数*右边数，先计算左边数，然后在计算右边数，循环两次
func productExceptSelf1(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}

	tmp := 1
	for i := len(nums) - 2; i >= 0; i-- {
		tmp = tmp * nums[i+1]
		ans[i] = ans[i] * tmp
	}
	return ans
}
