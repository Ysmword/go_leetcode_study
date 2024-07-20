package main

import (
	"fmt"
	"sort"
)

// 暴力算法：转map+两层遍历分类（加标志）
func groupAnagrams(strs []string) [][]string {
	// 转map
	strsMap := make([]map[byte]int, len(strs))
	for i := range strs {
		strsMap[i] = strMap(strs[i])
	}

	// 两层遍历分类
	visited := make([]int, len(strs))
	result := make([][]string, 0)
	for i := 0; i < len(strs); i++ {
		if visited[i] != 0 {
			continue
		}
		same := make([]string, 0)
		same = append(same, strs[i])
		visited[i] = 1
		for j := i + 1; j < len(strs); j++ {
			if visited[j] != 0 {
				continue
			}
			if isAnagram(strsMap[i], strsMap[j]) {
				same = append(same, strs[j])
				visited[j] = 1
			}
		}
		result = append(result, same)
	}
	return result
}

func strMap(s string) map[byte]int {
	sMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
	}
	return sMap
}

func isAnagram(sMap map[byte]int, tMap map[byte]int) bool {
	if len(sMap) != len(tMap) {
		return false
	}
	for item, count := range sMap {
		if count != tMap[item] {
			return false
		}
	}
	return true
}

// 算法2: 重点依旧是如何找到同类型，可以排序后找规律
func groupAnagrams2(strs []string) [][]string {
	sameMap := make(map[string][]string)
	for i := range strs {
		s := []byte(strs[i])
		sort.Slice(s, func(i, j int) bool {
			return s[i] > s[j]
		})
		sortS := string(s)
		sameMap[sortS] = append(sameMap[sortS], strs[i])
	}

	res := make([][]string, 0)
	for _, ss := range sameMap {
		res = append(res, ss)
	}
	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
