package main

import "fmt"

// 哈希表+三次遍历; 不对，这样没有考虑到顺序
func isIsomorphic1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap, tMap := strMap(s), strMap(t)

	for i := 0; i < len(s); i++ {
		if sMap[s[i]] != tMap[t[i]] {
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

// 双射关系，
/*
第一张哈希表 s2t 以 s 中字符为键，映射至 t 的字符为值，第二张哈希表 t2s 以 t 中字符为键，
映射至 s 的字符为值。从左至右遍历两个字符串的字符，不断更新两张哈希表，如果出现冲突（即当前
下标 index 对应的字符 s[index] 已经存在映射且不为 t[index] 或当前下标 index 对应的字
符 t[index] 已经存在映射且不为 s[index]）时说明两个字符串无法构成同构，返回 false。
*/
func isIsomorphic(s, t string) bool {
	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)
	for i := range s {
		x, y := s[i], t[i]
		// 对应的是下标值，
		// 超级注意：每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序
		// 也就意味着不会出现：用第二位字符，映射第一位字符
		if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
			return false
		}
		s2t[x] = y
		t2s[y] = x
	}
	return true
}

func main() {
	// fmt.Println(isIsomorphic1("egg", "add"))
	// fmt.Println(isIsomorphic1("foo", "bar"))
	// fmt.Println(isIsomorphic1("paper", "title"))

	// fmt.Println(isIsomorphic("egg", "add"))
	// fmt.Println(isIsomorphic("foo", "bar"))
	// fmt.Println(isIsomorphic("paper", "title"))
	fmt.Println(isIsomorphic("bbbaaaba", "aaabbbba"))
}
