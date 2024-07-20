package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortStr(s string) string {
	ss := make([]string, 0)
	for _, single := range s {
		ss = append(ss, string(single))
	}
	sort.Strings(ss)
	return strings.Join(ss, "")
}

// 排序+包含关系：不行
func canConstruct1(ransomNote string, magazine string) bool {
	ransomNote = sortStr(ransomNote)
	magazine = sortStr(magazine)
	fmt.Println(ransomNote, magazine)
	return strings.Contains(magazine, ransomNote)
}

// 哈希表
func canConstruct(ransomNote string, magazine string) bool {
	magazineMap := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		magazineMap[magazine[i]]++
	}
	can := true
	for i := 0; i < len(ransomNote); i++ {
		if count, ok := magazineMap[ransomNote[i]]; !ok || count == 0 {
			can = false
		} else {
			magazineMap[ransomNote[i]]--
		}
	}
	return can
}

func main() {
	// fmt.Println(canConstruct1("a", "b"))
	// fmt.Println(canConstruct1("aa", "ab"))
	// fmt.Println(canConstruct1("aa", "aab"))
	// fmt.Println(canConstruct1("aab", "baa"))
	// fmt.Println(canConstruct1("bg", "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj"))

	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
	fmt.Println(canConstruct("aab", "baa"))
	fmt.Println(canConstruct("bg", "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj"))
}
