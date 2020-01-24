/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] 添加与搜索单词 - 数据结构设计
 */
package problems

import "fmt"

type WordDictionary struct {
	isEnd bool
	next  [26]*WordDictionary
}

/** Initialize your data structure here. */
func Constructor_211() WordDictionary {
	return WordDictionary{}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	for _, c := range word {
		if this.next[c-'a'] == nil {
			this.next[c-'a'] = new(WordDictionary)
		}
		this = this.next[c-'a']
	}
	this.isEnd = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	var match func(word string, node *WordDictionary) bool
	match = func(word string, node *WordDictionary) bool {
		if len(word) == 0 {
			return node.isEnd
		}
		alpha := word[0]
		if alpha == '.' {
			for _, v := range node.next {
				if v != nil && match(word[1:], v) {
					return true
				}
			}
			return false
		} else {
			if node.next[alpha-'a'] == nil {
				return false
			}
			return match(word[1:], node.next[alpha-'a'])
		}
	}
	return match(word, this)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

// type WordDictionary struct {
// 	m map[int][]string
// }

// /** Initialize your data structure here. */
// func Constructor_211() WordDictionary {
// 	wd := WordDictionary{}
// 	wd.m = make(map[int][]string)
// 	return wd
// }

// /** Adds a word into the data structure. */
// func (this *WordDictionary) AddWord(word string) {
// 	this.m[len(word)] = append(this.m[len(word)], word)
// }

// /** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
// func (this *WordDictionary) Search(word string) bool {
// 	wordSlice, ok := this.m[len(word)]
// 	if !ok {
// 		return false
// 	}
// 	for i := range wordSlice {
// 		flag := true
// 		for j := range word {
// 			if word[j] == '.' {
// 				continue
// 			}
// 			if wordSlice[i][j] != word[j] {
// 				flag = false
// 			}
// 		}
// 		if flag == true {
// 			return true
// 		}
// 	}
// 	return false
// }

// @lc code=start

// @lc code=end
func Solve211() {
	wd := Constructor_211()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")
	fmt.Println(wd.Search("pad"))
	fmt.Println(wd.Search("bad"))
	fmt.Println(wd.Search(".ad"))
	fmt.Println(wd.Search("b.."))
}
