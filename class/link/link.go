package link

import (
	"fmt"
)

// ListNode 链表节点的定义，以及创建链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// CreateListBySlice 根据切片创建链表
func CreateListBySlice(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	// 创建头节点
	head := &ListNode{}
	cur := head
	for _, v := range nums {

		// 1. 创建下一个节点
		// 2. 连接下一个节点
		cur.Next = &ListNode{Val: v}

		// 3. 移动到下一个节点
		cur = cur.Next
	}
	return head.Next
}

func PrintlnList(l *ListNode) {
	cur := l
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}

	fmt.Println()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
