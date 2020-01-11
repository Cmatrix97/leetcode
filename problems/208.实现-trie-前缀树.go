/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */
package problems

import "fmt"

// @lc code=start
type Trie struct {
	isEnd bool
	next  [26]*Trie
}

/** Initialize your data structure here. */
func Constructor_208() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, c := range word {
		if this.next[c-'a'] == nil {
			this.next[c-'a'] = new(Trie)
		}
		this = this.next[c-'a']
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, c := range word {
		this = this.next[c-'a']
		if this == nil {
			return false
		}
	}
	return this.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, c := range prefix {
		this = this.next[c-'a']
		if this == nil {
			return false
		}
	}
	return true
}

func (this *Trie) DeleteWord(word string) {
	if !this.Search(word) {
		return
	}
	var delete func(node *Trie, index int)
	delete = func(node *Trie, index int) {
		if index == len(word) {
			node.isEnd = false
		} else {
			delete(node.next[word[index]-'a'], index+1)
		}
		if node.isEnd {
			return
		}
		for _, v := range node.next {
			if v != nil {
				return
			}
		}
		node = nil
	}
	delete(this, 0)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
func Solve208() {
	obj := Constructor_208()
	obj.Insert("apple")
	fmt.Println(obj.Search("apple"))
	fmt.Println(obj.Search("app"))
	fmt.Println(obj.StartsWith("app"))
	obj.Insert("app")
	fmt.Println(obj.Search("app"))
	obj.DeleteWord("app")
	fmt.Println(obj.Search("app"))
	fmt.Println(obj.Search("apple"))
}
