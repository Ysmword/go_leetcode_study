package main

import (
	"fmt"
	"sort"
)

/*
对citations数组从大到小进行排序
从头开始遍历排序后的citations数组：到达索引i，表示有i+1篇论文，被引用了citations[i]次，此时h指数就是论文数和被引用数二者较小值
如：[3,0,6,1,5]从大到小排序后为[6,5,3,1,0]，从头开始遍历：

	i==0，此时有i+1=1篇论文，被引用次数为citations[i]==6，因此此时h系数取较小值1
	i==1，此时有i+1=2篇论文，被引用次数为citations[i]==5，因此此时h系数取较小值2
	i==2，此时有i+1=3篇论文，被引用次数为citations[i]==3，因此此时h系数取较小值3
	i==3，此时有i+1=4篇论文，被引用次数为citations[i]==1，因此此时h系数取较小值1
	i==4，此时有i+1=5篇论文，被引用次数为citations[i]==0，因此此时h系数取较小值0

因此取上述所有情况的最大值，就是最终要求的结果
复杂度分析：

	时间复杂度：排序时间复杂度为O(nlog(n))，遍历一遍数组时间复杂度为O(n)，因此时间复杂度为O(nlog(n))
	空间复杂度：使用了常数级的额外空间，因此空间复杂度为O(1)
*/
func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	res := 0
	for i := 0; i < len(citations); i++ {
		res = max(res, min(i+1, citations[i]))
	}
	return res
}

func main() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
	fmt.Println(hIndex([]int{1, 3, 1}))
}
