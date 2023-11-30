package link

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/add-two-numbers/
// 给两个链表，每个链表代表一个非负整数，每个节点代表一位数字，将两个数相加，返回一个新的链表
// 解法1
func addTwoNumbersV1(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	carry := 0
	cur, cur1, cur2 := dummy, list1, list2
	// 1. 遍历完最短的那个链表，处理两数相加以及进位
	for cur1 != nil && cur2 != nil {
		sum := cur1.Val + cur2.Val + carry
		carry = sum / 10

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	// 2. 遍历未完成的链表，处理两数相加以及进位
	if cur2 != nil {
		cur1 = cur2
	}

	for cur1 != nil {
		sum := cur1.Val + carry
		carry = sum / 10

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		cur1 = cur1.Next
	}

	if carry != 0 {
		cur.Next = &ListNode{Val: carry}
	}

	// 3. 返回相应的结果
	return dummy.Next
}

// 解法2
// 上述解法在处理的时候，考虑到了cur1, cur2为nil的情况，所以需要额外的处理
// 当时，整个算法最核心的就是，cur1.Val + cur2.Val + carry这个值的计算
// 这个计算为 最长的那个链表， 以及两者的和，以及进位
func addTwoNumbersV2(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur, cur1, cur2 := dummy, list1, list2
	carry := 0

	for cur1 != nil || cur2 != nil || carry != 0 {
		// v1， v2表示每次遍历的值，作用域应该放for里面
		v1, v2 := 0, 0
		if cur1 != nil {
			v1 = cur1.Val
			cur1 = cur1.Next
		}

		if cur2 != nil {
			v2 = cur2.Val
			cur2 = cur2.Next
		}

		sum := v1 + v2 + carry
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return dummy.Next
}

func TestAddTwoNumbers(t *testing.T) {
	list1 := CreateListBySlice([]int{1, 2, 3, 4, 5})
	list2 := CreateListBySlice([]int{1, 2, 3, 4, 5})

	list3 := addTwoNumbersV1(list1, list2)
	list4 := addTwoNumbersV2(list1, list2)

	PrintlnList(list3)
	PrintlnList(list4)
}

// 删除链表的倒数第N个节点
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	// 第一性原则，要删除倒数第n个节点，那就需要找到倒数第n+1个节点；

	// 因为，我们采用快慢指针的方式，那么需要考虑，快指针什么时候结束，当快指针移动到nil指针的时候，结束；
	// 快指针，先移动n+1下，然后慢指针才开始移动；
	// 注意着几个对应关系；
	// fast到nil，终止；如果两个节点相差1，表示慢节点就是倒数第一个节点，因为我们要找倒数第n+1节点，所以快指针先移动n+1下，然后慢指针才开始移动；
	// 最后处理链接关系；
	dummmy := &ListNode{}
	dummmy.Next = head
	fast, slow := dummmy, dummmy
	for i := 0; i < n+1 && fast != nil; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return dummmy.Next
}

func TestRemoveNthFromEnd(t *testing.T) {
	list := CreateListBySlice([]int{1, 2, 3, 4, 5})
	list = removeNthFromEnd(list, 2)

	PrintlnList(list)

	list = removeNthFromEnd(list, 1)
	PrintlnList(list)

	list = removeNthFromEnd(list, 3)
	PrintlnList(list)
}

// 21. 合并两个有序链表
// https://leetcode.com/problems/merge-two-sorted-lists/
// 解法1：迭代方法
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}

		cur = cur.Next
	}

	if l1 != nil {
		l2 = l1
	}

	for l2 != nil {
		cur.Next = l2
		l2 = l2.Next
		cur = cur.Next
	}

	return dummy.Next
}

// 解法2：递归方法
func mergeTwoListsV2(l1 *ListNode, l2 *ListNode) *ListNode {
	// 终止条件
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	// 处理本层的逻辑以及与下一层的连接
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}

	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2

}
func TestMergeTwoLists(t *testing.T) {
	l1 := CreateListBySlice([]int{1, 2, 3, 4, 5})
	l2 := CreateListBySlice([]int{1, 2, 3, 4, 5})
	PrintlnList(mergeTwoLists(l1, l2))

	l3 := CreateListBySlice([]int{1, 2, 3, 4, 5})
	l4 := CreateListBySlice([]int{1, 2, 3, 4, 5})
	PrintlnList(mergeTwoListsV2(l3, l4))
}

// 23. 合并K个升序链表
// https://leetcode-cn.com/problems/merge-k-sorted-lists/
// 解法1：分治法: 先将两个进行合并，直到合并到只有一个链接为止
// 解法2：优先队列
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) <= 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	mid := len(lists) / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	return mergeTwoListsV2(left, right)
}

// 优先队列
func mergeKListsV2(lists []*ListNode) *ListNode {

	return nil
}

// 24. Swap Nodes in Pairs
// https://leetcode.com/problems/swap-nodes-in-pairs/
// 解法1：迭代法
// 这里应该如何想呢？我需要交换两个节点，那么一定会设计到3个节点；
// 将这个三个节点为一个整天，那么，我们首先要断链，将这个整体最后一个节点的next保存起来；
// 接着处理每一个元素，比如这里，我们先处理pre，在处理最后一个，继续处理，直到把所有元素处理完；
// 处理完这个整体之后，处理下一个整体；
func swapPairs(head *ListNode) *ListNode {
	// 建立哨兵节点
	dummy := &ListNode{Next: head}
	pre, cur := dummy, head

	// 现在是两两交换，因此设计到3个节点，我们将3个节点作为一个整体，每次处理3个节点
	for cur != nil && cur.Next != nil {
		// 断链以及一些初始化操作
		end := cur.Next.Next
		first := cur
		second := cur.Next

		// 处理交换的流程
		pre.Next = second
		second.Next = first
		first.Next = end

		// 处理下一个整体
		pre = first
		cur = first.Next
	}

	return dummy.Next
}

func Test_SwapPairs(t *testing.T) {
	list := CreateListBySlice([]int{1, 2, 3, 4, 5})
	res := swapPairs(list)
	PrintlnList(res)
}

func TestSwapK(t *testing.T) {
	list := CreateListBySlice([]int{1, 2, 3, 4, 5})

	newhead := swapK(list, 3)
	PrintlnList(newhead)
}

// 交换k个元素
// 我们需要看看，写出这道题，需要哪些信息呢？
// https://leetcode.com/problems/reverse-nodes-in-k-group/
func swapK(head *ListNode, k int) *ListNode {
	// 使用递归来求解
	// 递归的终止条件
	cur := head
	// 截取k个元素为一个整体
	for i := 0; i < k; i++ {
		// 少于k个元素，不进行反转
		if cur == nil {
			return head
		}

		// 移动k次，此时到达，下一个整体的第一个元素
		cur = cur.Next
	}

	newHead := reverse(head, cur)
	head.Next = swapK(cur, k)

	return newHead
}

// 给定一个区间，将这个区间进行反转
func reverse(first *ListNode, last *ListNode) *ListNode {
	// 反转之后，第一个元素应该链接到链表的最后一个元素
	pre := last
	for first != last {
		// 保留下一个节点
		next := first.Next

		// 移动当前节点
		first.Next = pre
		pre = first
		first = next
	}

	return pre
}

// 链表旋转
// https://leetcode.com/problems/rotate-list/
// 题目简介
// 给定一个链表，将链表向右旋转k个位置，其中k是非负数。
func rotateRight(head *ListNode, k int) *ListNode {
	// 1. 先将链表闭合成环
	// 2. 找到相应的位置，断开环
	// 3. 返回新的头节点
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// 1. 我们需要到达最后一个节点，并且计算链表长度，以及将链表闭合成环
	// 假设，我们有n个节点，那么，for里面的代码应该执行（n-1）次；
	cur := head
	n := 1
	for cur.Next != nil {
		cur = cur.Next
		n++
	}

	// 链接成一个环
	cur.Next = head

	// 2. 找到相应的位置，断开环， 要找到倒数第k+1个元素；总共n个元素，也就是正向的 n - k - 1 个元素；
	cur = head
	// 它到的位置为，n - k % n
	for i := 0; i < n-k%n-1; i++ {
		cur = cur.Next
	}

	// 找到新
	newHead := cur.Next
	cur.Next = nil

	return newHead
}

// 统计链表的长度
func countList(head *ListNode) int {
	cur := head
	n := 0

	// 多少个节点，执行for循环逻辑多少次
	// 以终为始：我们的目标，是获取链表中有多少个节点，那么我们希望，每个节点都执行+1操作
	for cur != nil {
		cur = cur.Next
		n++
	}

	return n
}

// "以终为始"是一个汉语成语，直译为“starting with the end in mind”。
// 这个成语的含义是在做任何事情之前，首先要明确最终的目标或结果，然后再计划和执行过程。
// 它强调目标导向和前瞻性思维的重要性，提醒人们在做决策或规划时要考虑长远的影响和结果。
// 简而言之，这是一种以目标为导向的规划和行动方式。
func TestCountList(t *testing.T) {
	list := CreateListBySlice([]int{1, 2, 3, 4, 5})
	cnt := countList(list)
	fmt.Println(cnt)
}

// 移动到链表最后一个节点
func TestLastNode(t *testing.T) {
	list := CreateListBySlice([]int{1, 2, 3, 4, 5})
	cur := list
	for cur != nil && cur.Next != nil {
		cur = cur.Next
	}

	fmt.Println(cur.Val)
}

// 链表的倒数第k个节点
func TestLastKNode(t *testing.T) {
	list := CreateListBySlice([]int{1, 2, 3, 4, 5})
	node := LaskK(list, 2)
	fmt.Println(node.Val)
}

// 倒数第k个节点
func LaskK(list *ListNode, k int) *ListNode {
	// 保证倒数第第k个节点一定存在
	// 移动k-1次，即能到达倒数第k个节点
	// 要到达地k个节点，就只需要移动k-1次即可
	fast, slow := list, list
	for i := 0; i < k-1; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}

func TestLaskK1(t *testing.T) {
	list := CreateListBySlice([]int{1, 2, 3, 4, 5})
	node := LaskK1(list, 2)

	fmt.Println(node.Val)
}

// 但是，如果我们需要删除第k个节点，那么就需要找到倒数第k+1个节点
// 要找到倒数第k+1个几点，首先，快节点，应该移动到第K+1个节点
func LaskK1(list *ListNode, k int) *ListNode {
	fast, slow := list, list
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}

// 为什么需要头结点呢？
// 如果倒数第k个节点是第一个节点的情况，就找不到倒数第k+1个节点了
// 为了解决这个问题，我们需要引入头结点
// 删除倒数第k个节点
func deleteLastKNode(list *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = list

	fast, slow := dummy, dummy
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}

func TestDeleteLastK(t *testing.T) {
	list := CreateListBySlice([]int{1, 2, 3, 4, 5})
	node := deleteLastKNode(list, 2)
	PrintlnList(node)
}
