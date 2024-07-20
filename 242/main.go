package main

import "fmt"

// 直接哈希表计算各个字符出现次数即可
func isAnagram(s string, t string) bool {
	sMap, tMap := strMap(s), strMap(t)
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

func strMap(s string) map[byte]int {
	sMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
	}
	return sMap
}

// 另外的算法：排序后比较，用数组

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "cat"))
}
