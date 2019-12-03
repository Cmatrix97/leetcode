package problem

//Solution167 两数之和II-输入有序数组
func Solution167()  {
	// numbers := []int{2,7,11,15}
	// target := 9
	numbers := []int{2, 3, 4}
	target := 6
	for _, v := range twoSum(numbers, target) {
		print(v)
	}
}

/**利用map
1.遍历时向map中存储key=当前数字,value=坐标
2.同时查询map中是否有key=target-当前数字,如果有，当前的遍历的index和map[key]即所求
*/
func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for i, v := range numbers {
		if v, ok := m[target - v]; ok {
			return []int{v+1, i+1}
		}
		m[v] = i
	}
	return []int{}
}