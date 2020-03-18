package interview

import "fmt"

// 利用map去重
func permutation1(s string) []string {
	visited := make([]bool, len(s))
	var res []string
	var tmp []byte
	m := make(map[string]struct{})
	var recurse func()
	recurse = func() {
		if len(tmp) == len(s) {
			str := string(tmp)
			if _, ok := m[str]; !ok {
				m[str] = struct{}{}
				res = append(res, str)
			}
			return
		}
		for i := range s {
			if visited[i] {
				continue
			}
			visited[i] = true
			tmp = append(tmp, s[i])
			recurse()
			visited[i] = false
			tmp = tmp[:len(tmp)-1]
		}
	}
	recurse()
	return res
}

// 交换法
func permutation2(s string) []string {
	var res []string
	tmp := []byte(s)
	var helper func(i int)
	helper = func(i int) {
		if i == len(tmp) {
			res = append(res, string(tmp))
		}
		seen := make(map[byte]bool)
		for j := i; j < len(tmp); j++ {
			if seen[tmp[j]] {
				continue
			}
			tmp[i], tmp[j] = tmp[j], tmp[i]
			helper(i + 1)
			tmp[i], tmp[j] = tmp[j], tmp[i]
			seen[tmp[j]] = true
		}
	}
	helper(0)
	return res
}

func SolveOffer38() {
	s := "aab"
	fmt.Println(permutation1(s))
}
