package main

import "fmt"

// 方法1：判断买入的第二天价格比第一天的小，就不买入；只有小的时候才买入，并且在第二天就卖出
// 主要思想：画折线图，就可以看到只要买升，你得到的利润永远是最大的
func maxProfit(prices []int) int {
	max := 0
	for i := 0; i < len(prices); i++ {
		if i == len(prices)-1 {
			continue
		}
		if prices[i] < prices[i+1] {
			max += prices[i+1] - prices[i]
		}
	}
	return max
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
