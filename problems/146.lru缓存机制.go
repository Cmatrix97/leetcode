/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 */
package problems

import "fmt"

type LRUNode struct {
	key, value int
	prev, next *LRUNode
}

type LRUCache struct {
	head, tail *LRUNode
	m          map[int]*LRUNode
	cap        int
}

func Constructor146(capacity int) LRUCache {
	var lruc LRUCache
	lruc.head, lruc.tail = new(LRUNode), new(LRUNode)
	lruc.head.next, lruc.tail.prev = lruc.tail, lruc.head
	lruc.m = make(map[int]*LRUNode)
	lruc.cap = capacity
	return lruc
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		this.removeNode(node)
		this.insertToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		this.removeNode(node)
		this.insertToHead(node)
		node.value = value
		return
	}

	if len(this.m) == this.cap {
		delete(this.m, this.tail.prev.key)
		this.removeNode(this.tail.prev)
	}

	node := new(LRUNode)
	node.key, node.value = key, value
	this.insertToHead(node)
	this.m[key] = node
}

func (this *LRUCache) removeNode(node *LRUNode) {
	node.prev.next, node.next.prev = node.next, node.prev
}

func (this *LRUCache) insertToHead(node *LRUNode) {
	this.head.next, this.head.next.prev, node.prev, node.next = node, node, this.head, this.head.next
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=start

// @lc code=end
func Solve146() {
	cache := Constructor146(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}
