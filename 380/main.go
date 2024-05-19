package main

import (
	"fmt"
	"math/rand"
)

// 方法：变长数组 + 哈希表
type RandomizedSet struct {
	m map[int]int // key is value, v is index
	s []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m: make(map[int]int),
		s: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; !ok {
		this.s = append(this.s, val)
		this.m[val] = len(this.s) - 1
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.m[val]
	if !ok {
		return false
	}
	delete(this.m, val)
	if len(this.s) <= 1 {
		this.s = make([]int, 0)
	} else if index == len(this.s)-1 {
		this.s = this.s[:len(this.s)-1]
	} else {
		this.s[index], this.s[len(this.s)-1] = this.s[len(this.s)-1], this.s[index]
		this.s = this.s[:len(this.s)-1]
		this.m[this.s[index]] = index
	}
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.s[rand.Intn(len(this.s))]
}

func main() {
	r := Constructor()
	// fmt.Println(r.Insert(1))
	// fmt.Println(r.Remove(2))
	// fmt.Println(r.Insert(2))
	// fmt.Println(r.GetRandom())
	// fmt.Println(r.Remove(1))
	// fmt.Println(r.Insert(2))
	// fmt.Println(r.GetRandom())
	fmt.Println(r.Insert(3))
	fmt.Println(r.Insert(3))
	fmt.Println(r.GetRandom())
	fmt.Println(r.GetRandom())
	fmt.Println(r.Insert(1))
	fmt.Println(r.Remove(3))
	fmt.Println(r.GetRandom())
	fmt.Println(r.GetRandom())
	fmt.Println(r.Insert(0))
	fmt.Println("r:", r)
	fmt.Println(r.Remove(0))
}
