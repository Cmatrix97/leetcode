package problem

//Solution26 删除排序数组中的重复项
func Solution26()  {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	length := removeDuplicates(nums)
	println(length)
	for _, v := range nums {
		print(v)
	}
}

/*快慢双指针
1.开始时都指向第二个数字
2.(1)如果快指针所指数字与前一个数字相同，则快指针向前走一步
    (2)如果不同，则两个指针都向前走一步
3.当快指针遍历整个数组后，慢指针当前的坐标即新数组长度
*/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		nums[j] = nums[i]
		j++
	}
	return j
}