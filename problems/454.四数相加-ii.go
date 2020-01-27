/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */
package problems

import "fmt"

/*
暴力求解
O(n^4)
超时31/48
*/
func fourSumCount1(A []int, B []int, C []int, D []int) int {
	var count int
	for _, a := range A {
		for _, b := range B {
			for _, c := range C {
				for _, d := range D {
					if a+b+c+d == 0 {
						count++
					}
				}
			}
		}
	}
	return count
}

/*
两个map分别存储a+b和c+d
超时46/48
*/
func fourSumCount2(A []int, B []int, C []int, D []int) int {
	var count int
	mapAB, mapCD := make(map[int]int), make(map[int]int)
	for _, a := range A {
		for _, b := range B {
			mapAB[a+b]++
		}
	}
	for _, c := range C {
		for _, d := range D {
			mapCD[c+d]++
		}
	}
	fmt.Println(mapAB, mapCD)
	for ab, c1 := range mapAB {
		for cd, c2 := range mapCD {
			if ab+cd == 0 {
				count += c1 * c2
			}
		}
	}
	return count
}

/*
一个map存储A+B
遍历C+D时直接查找
*/
func fourSumCount3(A []int, B []int, C []int, D []int) int {
	var count int
	mapAB := make(map[int]int)
	for _, a := range A {
		for _, b := range B {
			mapAB[a+b]++
		}
	}
	for _, c := range C {
		for _, d := range D {
			count += mapAB[-(c + d)]
		}
	}
	return count
}

// @lc code=start

// @lc code=end
func Solve454() {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}
	fmt.Println(fourSumCount1(A, B, C, D))
}
