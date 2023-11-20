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
