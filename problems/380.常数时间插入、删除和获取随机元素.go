/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] 常数时间插入、删除和获取随机元素
 */
package problems

import (
	"fmt"
	"math/rand"
)

// @lc code=start
type RandomizedSet struct {
	m    map[int]int
	nums []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	rs := RandomizedSet{}
	rs.m = make(map[int]int)
	return rs
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.m[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	valIndex, ok := this.m[val]
	if !ok {
		return false
	}
	lastIndex := len(this.nums) - 1
	this.m[this.nums[lastIndex]] = valIndex
	this.nums[valIndex] = this.nums[lastIndex]
	this.nums = this.nums[:lastIndex]
	delete(this.m, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.nums))
	return this.nums[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
func Solve380() {
	randomSet := Constructor()
	fmt.Println(randomSet.Insert(1))
	fmt.Println(randomSet.Remove(2))
	fmt.Println(randomSet.Insert(2))
	fmt.Println(randomSet.GetRandom())
	fmt.Println(randomSet.Remove(1))
	fmt.Println(randomSet.Insert(2))
	fmt.Println(randomSet.GetRandom())
	// true false true rand(1,2) true false 2
}
