package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	data := make([]int, 0)
	for x != 0 {
		data = append(data, x%10)
		x = x / 10
	}
	for i := 0; i < len(data)/2; i++ {
		if data[i] != data[len(data)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(121))
}
