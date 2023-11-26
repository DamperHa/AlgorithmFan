package link

import (
	"fmt"
	"testing"
)

// 写题的作用是什么呢？
//  我们需要好好捋一捋，万物之间都有联系，可以将生活中的一些方法，和题目之间进行类比
//  刷题也是面试过程中，重要的一环，有这一条，你就没有不刷的理由
//  明白什么是LRU，Least Recently Used, 最近最少使用算法，如果容器不够，当新增某个节点的时候，需要将最久没使用的那个节点删除

// 为了使用O(1)的算法复杂度，我们使用map + List 来实现；
type LRUCache struct {
	keys       map[int]*Node
	head, tail *Node
	cap        int
}

type Node struct {
	Val, Key  int
	Pre, Next *Node
}

func Constructor(capacity int) LRUCache {
	// 定义一个头结点嗯？
	headDummy := &Node{}
	tailDumy := &Node{}

	headDummy.Next = tailDumy
	tailDumy.Pre = headDummy

	return LRUCache{
		keys: make(map[int]*Node),
		cap:  capacity,
		tail: tailDumy,
		head: headDummy,
	}
}

//	如果key存在LRU中，将该节点移动到链表首部，并返回对应的值；
//
// 将节点移动到首部，包含两个方面，一个是在链表中删除该节点，一个是将该节点添加到首部；
//
//	如果key不存在LRU中，直接返回-1
func (this *LRUCache) Get(key int) int {

	if node, ok := this.keys[key]; ok {
		this.Remove(node)
		this.Add(node)

		return node.Val
	}

	return -1
}

// 如果在链表中，更新节点，移动到首部；
// 如果不在链表，添加到链表首部；
// 最后加测链表长度是否超过最大长度；
// 【这里写法很有意思，我们先执行某个操作，然后再检测LRU状态是否符合要求】
func (this *LRUCache) Put(key int, value int) {
	//  存在
	if node, ok := this.keys[key]; ok {
		node.Val = value
		this.Remove(node)
		this.Add(node)

		return
	} else {
		node = &Node{Val: value, Key: key}
		this.keys[key] = node
		this.Add(node)
	}

	if len(this.keys) > this.cap {
		delete(this.keys, this.tail.Pre.Key)
		this.Remove(this.tail.Pre)

	}
}

// 从上面的操作，我提取出两个方法，一个是删除给定节点，一个是在头结点增加一个节点
// 双向链表如何处理？对于这种以前解除比较少的，我们如何更快的掌握呢？抽象与总结
//
//	这里不管怎么样，你还是要区分，是第一个节点，还是最后一个节点
func (this *LRUCache) Add(node *Node) {
	node.Next = this.head.Next
	node.Pre = this.head

	this.head.Next = node
	node.Next.Pre = node
}

// 删除给定节点
// 处理两个事情，一个前驱节点的后续节点，一个是后续节点的前驱指针
func (this *LRUCache) Remove(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func PrintlnList2(l *Node) {
	cur := l
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}

	fmt.Println()
}

func TestLRU(t *testing.T) {
	lru := Constructor(2)

	lru.Put(1, 1)
	PrintlnList2(lru.head)

	lru.Put(2, 2)
	PrintlnList2(lru.head)

	lru.Get(1)
	PrintlnList2(lru.head)

	lru.Put(3, 3)

	PrintlnList2(lru.head)

	lru.Get(2)
	fmt.Println("+++®")
	PrintlnList2(lru.head)
	lru.Put(4, 4)

	fmt.Println("444444444444444")
	PrintlnList2(lru.head)

	lru.Get(1)
	lru.Get(3)
	lru.Get(4)

}

func TestStr(t *testing.T) {
	str := "fanzhihao"
	fmt.Printf("%T", str[1])

	for _, r := range str {
		fmt.Printf("%T", r)
	}
}

// LFU
// LRU中使用的是双向链表，在LFU中还是使用双向链表，将相同频率的凡在一个链表，并且每次放在末尾
//一个链表，``	-破了批评了；
// 查找元素：删除对应节点，将节点添加到首部；
// 增加元素：链表头部；
// 超出容量：删除元素，链表尾部；

// 查找：在对应fre删除node，并在新的fre的末尾添加元素；
// 增加元素：如果不存在，在1的链表添加元素，如果存在，删除原fre节点， 在新fre添加节点；
// 超出容量：删除最小fre的第一个元素；（从插入与获取有一定的区别）
