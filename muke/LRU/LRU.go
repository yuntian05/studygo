package main

import "fmt"

type Node struct {
	Key   int
	Value int
	Pre   *Node
	Next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

// 新建双向链表
func NewList() *LinkedList {
	list := &LinkedList{}
	list.head = &Node{
		Key:   0,
		Value: 0,
	}
	list.tail = &Node{
		Key:   0,
		Value: 0,
	}
	list.head.Next = list.tail
	list.tail.Pre = list.head
	list.len = 0
	return list
}

func (list *LinkedList) Len() int {
	return list.len
}

// 在链表尾部添加节点 x，时间 O(1)
func (list *LinkedList) addLastNode(node *Node) {
	node.Pre = list.tail.Pre
	node.Next = list.tail
	list.tail.Pre.Next = node
	list.tail.Pre = node
	list.len += 1
}

// 删除链表中的 x 节点（x 一定存在）
// 由于是双链表且给的是目标 Node 节点，时间 O(1)
func (list *LinkedList) remove(node *Node) {
	pre := node.Pre
	next := node.Next
	pre.Next = next
	next.Pre = pre
	list.len -= 1
}

// 删除链表中第一个节点，并返回该节点，时间 O(1)
func (list *LinkedList) removeFirst() *Node {
	if list.head.Next == nil {
		return nil
	}
	first := list.head.Next
	list.remove(first)
	return first
}

type LRUCache struct {
	m     map[int]*Node
	cache *LinkedList
	cap   int
}

func NewLRU(capacity int) *LRUCache {
	lru := &LRUCache{}
	lru.m = make(map[int]*Node)
	lru.cache = NewList()
	lru.cap = capacity
	return lru
}

//将某个 key 提升为最近使用的
func (lru *LRUCache) makeRecently(key int) {
	node, ok := lru.m[key]
	if !ok {
		return
	}
	// 先从链表中删除该节点
	lru.cache.remove(node)
	// 直接插入到队尾
	lru.cache.addLastNode(node)
}

//添加最近使用的元素
func (lru *LRUCache) addRecently(key, val int) {
	node := &Node{Value: val, Key: key}
	// 链表尾部就是最近使用的元素
	lru.cache.addLastNode(node)
	// 别忘了在 map 中添加 key 的映射
	lru.m[key] = node
}

// 删除某一个 key
func (lru *LRUCache) deleteKey(key int) {
	node, ok := lru.m[key]
	if !ok {
		return
	}
	// 从链表中删除
	lru.cache.remove(node)
	// 从 map 中删除
	delete(lru.m, key)
}

func (lru *LRUCache) removeLeastRecently() {
	// 链表头部的第一个元素就是最久未使用的
	deletedNode := lru.cache.removeFirst()
	if deletedNode == nil {
		return
	}
	// 同时别忘了从 map 中删除它的 key
	key := deletedNode.Key
	delete(lru.m, key)
}

func (lru *LRUCache) Get(key int) int {
	node, ok := lru.m[key]
	if !ok {
		return -1
	}
	lru.makeRecently(key)
	return node.Value
}

func (lru *LRUCache) Put(key, val int) {
	_, ok := lru.m[key]
	if ok {
		// 删除旧的数据
		lru.deleteKey(key)
		// 新插入的数据为最近使用的数据
		lru.addRecently(key, val)
		return
	}
	// 容量满了
	if lru.cap == lru.cache.Len() {
		// 删除最久未使用的元素
		lru.removeLeastRecently()
	}
	// 添加为最近使用的元素
	lru.addRecently(key, val)
}

func main(){
	lru := NewLRU(3)
	lru.Put(1,1)
	lru.Put(2,2)
	lru.Put(3,3)
	fmt.Println(lru.Get(1))
	lru.Put(4,4)
	fmt.Println(lru.Get(2))

}
