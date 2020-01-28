/*
 * @lc app=leetcode.cn id=609 lang=golang
 *
 * [609] 在系统中查找重复文件
 */
package problems

import (
	"fmt"
	"strings"
)

func findDuplicate1(paths []string) [][]string {
	m := make(map[string][]string)
	for _, path := range paths {
		values := strings.Split(path, " ")
		pre := values[0]
		for i := 1; i < len(values); i++ {
			idx := strings.Index(values[i], "(")
			post := values[i][:idx]
			content := values[i][idx:]
			m[content] = append(m[content], pre+"/"+post)
		}
	}
	var res [][]string
	for _, v := range m {
		if len(v) > 1 {
			res = append(res, v)
		}
	}
	return res
}

// @lc code=start

// @lc code=end
func Solve609() {
	paths := []string{
		"root/a 1.txt(abcd) 2.txt(efgh)",
		"root/c 3.txt(abcd)",
		"root/c/d 4.txt(efgh)",
		"root 4.txt(efgh)",
	}
	res := findDuplicate1(paths)
	fmt.Println(res)
}
