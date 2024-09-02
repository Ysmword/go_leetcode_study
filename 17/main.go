package main

import "fmt"

var (
	dm = map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
)

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	res := make([]string, 0)
	backtack(0, digits, "", &res)
	return res
}

func backtack(start int, digits string, track string, res *[]string) {
	if start <= len(digits) { // 不要忘了等于号，因为要加入结果集
		if len(track) == len(digits) {
			// 拷贝一份
			track2 := make([]byte, len(track))
			copy(track2, track)
			*res = append(*res, string(track2))
			return
		}
		for i := 0; i < len(dm[string(digits[start])]); i++ {
			track += dm[string(digits[start])][i]
			backtack(start+1, digits, track, res)
			track = track[:len(track)-1]
		}
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
}
