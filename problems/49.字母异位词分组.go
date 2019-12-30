/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */
package problems

import (
	"fmt"
)

/*
将26个字母的计数数组作为map的key,value记录res中行下标
*/
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	mark := make(map[[26]byte]int)
	for _, str := range strs {
		var key [26]byte
		for _, r := range str {
			key[r-'a']++
		}
		if i, ok := mark[key]; !ok {
			res = append(res, []string{str})
			mark[key] = len(mark)
		} else {
			res[i] = append(res[i], str)
		}
	}
	return res
}

// @lc code=start

// @lc code=end

func Solve49() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	fmt.Println(res)
}
