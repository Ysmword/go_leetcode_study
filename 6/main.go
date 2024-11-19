package main

import (
	"fmt"
	"strings"
)

// 将每一行的字符拼接起来，最后拼接起来
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]string, numRows)
	row, flag := 0, -1
	for i := 0; i < len(s); i++ {
		rows[row] += string(s[i])
		if row == 0 || row == numRows-1 {
			flag = -flag // 等于零时，取反表示向下，等于numRows-1时，取反表示向上
		}
		row += flag
	}
	return strings.Join(rows, "")
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
