package interview

import "fmt"

// 暴力法 超时
func nthUglyNumber1(n int) int {
	ans := 1
	for {
		cur := ans
		for cur%2 == 0 {
			cur /= 2
		}
		for cur%3 == 0 {
			cur /= 3
		}
		for cur%5 == 0 {
			cur /= 5
		}
		if cur == 1 {
			n--
		}
		if n == 0 {
			break
		}
		ans++
	}
	return ans
}

// 三指针
func nthUglyNumber2(n int) int {
	var arr [1690]int
	arr[0] = 1
	var p2, p3, p5 int
	for i := 1; i < n; i++ {
		v2, v3, v5 := arr[p2]*2, arr[p3]*3, arr[p5]*5
		val := v2
		if v3 < val {
			val = v3
		}
		if v5 < val {
			val = v5
		}
		if val == v2 {
			p2++
		}
		if val == v3 {
			p3++
		}
		if val == v5 {
			p5++
		}
		arr[i] = val
	}
	return arr[n-1]
}

func SolveOffer49() {
	n := 1500
	fmt.Println(nthUglyNumber2(n))
}
