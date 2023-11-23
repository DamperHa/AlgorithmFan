package link

import (
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
