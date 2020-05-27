/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU缓存
 */
package problems

import "fmt"

// 链表节点 包含访问的次数
type LFUNode struct {
	key, value, freq int
	prev, next       *LFUNode
}

// 双链表
type DoubleLinkList struct {
	head, tail *LFUNode
}

type LFUCache struct {
	m       map[int]*LFUNode
	freqs   map[int]*DoubleLinkList
	cap     int
	minFreq int
}

func Constructor460(capacity int) LFUCache {
	var lfuc LFUCache
	lfuc.cap = capacity
	lfuc.m = make(map[int]*LFUNode)
	lfuc.freqs = make(map[int]*DoubleLinkList)
	return lfuc
}

func (this *LFUCache) Get(key int) int {
	if this.cap == 0 {
		return -1
	}

	if node, ok := this.m[key]; ok {
		this.access(node)
		return node.value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}

	// 如果存在修改值
	if node, ok := this.m[key]; ok {
		this.access(node)
		node.value = value
		return
	}

	// 不存在且cap已满
	if len(this.m) == this.cap {
		dll := this.freqs[this.minFreq]
		delNode := dll.tail.prev
		this.removeNode(delNode)
		delete(this.m, delNode.key)
	}

	node := new(LFUNode)
	node.key, node.value, node.freq = key, value, 1
	this.insertToHead(node)
	this.m[key] = node
	this.minFreq = 1
}

func (this *LFUCache) access(node *LFUNode) {
	// 从原来位置移除
	this.removeNode(node)
	// 添加到freq+1的位置
	node.freq++
	this.insertToHead(node)
}

func (this *LFUCache) initDoubleLinkList(freq int) {
	if _, ok := this.freqs[freq]; ok {
		return
	}
	dll := new(DoubleLinkList)
	dll.head, dll.tail = new(LFUNode), new(LFUNode)
	dll.head.next, dll.tail.prev = dll.tail, dll.head
	this.freqs[freq] = dll
}

func (this *LFUCache) insertToHead(node *LFUNode) {
	this.initDoubleLinkList(node.freq)
	dll := this.freqs[node.freq]
	dll.head.next, dll.head.next.prev, node.prev, node.next = node, node, dll.head, dll.head.next
	this.freqs[node.freq] = dll
}

func (this *LFUCache) removeNode(node *LFUNode) {
	dll := this.freqs[node.freq]
	node.prev.next, node.next.prev = node.next, node.prev
	if node.freq == this.minFreq && dll.head.next == dll.tail {
		this.minFreq++
	}
	this.freqs[node.freq] = dll
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=start

// @lc code=end
func Solve460() {
	cache := Constructor460(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(3))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}
