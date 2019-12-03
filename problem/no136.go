package problem

//Solution136 只出现一次的数字
func Solution136()  {
	nums := []int{4,1,2,1,2}
	println(singleNumber(nums))
}

/**异或运算
1.两个相同的数异或为0       t ^ t = 0
2.0异或一个数等于那个数  0 ^ t = t
*/
func singleNumber(nums []int) int {
	t := 0
	for _, v := range nums {
		t ^= v
	}
	return t
}