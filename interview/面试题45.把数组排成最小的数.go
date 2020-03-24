package interview

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

// 实现sort.Interface接口
func minNumber1(nums []int) string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}
	sort.Sort(sortBy1(strs))
	var buf bytes.Buffer
	for _, str := range strs {
		buf.WriteString(str)
	}
	return buf.String()
}

type sortBy1 []string

func (s sortBy1) Len() int {
	return len(s)
}

func (s sortBy1) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortBy1) Less(i, j int) bool {
	return solve(s[i], s[j])
}

func solve(s1, s2 string) bool {
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] < s2[i] {
			return true
		} else if s1[i] > s2[i] {
			return false
		}
	}
	if len(s1) > len(s2) {
		return solve(s1[len(s2):], s2)
	} else if len(s1) < len(s2) {
		return solve(s1, s2[len(s1):])
	} else {
		return true
	}
}

func minNumber2(nums []int) string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}
	sort.Sort(sortBy2(strs))
	var buf bytes.Buffer
	for _, str := range strs {
		buf.WriteString(str)
	}
	return buf.String()
}

type sortBy2 []string

func (s sortBy2) Len() int {
	return len(s)
}

func (s sortBy2) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortBy2) Less(i, j int) bool {
	a, _ := strconv.Atoi(s[i] + s[j])
	b, _ := strconv.Atoi(s[j] + s[i])
	return a < b
}

func SolveOffer45() {
	nums := []int{3, 30, 34, 5, 9}
	fmt.Println(minNumber1(nums))
}
