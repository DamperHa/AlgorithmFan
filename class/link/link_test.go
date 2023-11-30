package link

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateListBySlice(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateListBySlice(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateListBySlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	// 反转slow后面所有的节点
	PrintlnList(slow)
	var pre *ListNode
	for slow != nil {
		nxt := slow.Next

		slow.Next = pre
		pre = slow
		slow = nxt
	}

	PrintlnList(pre)
	cur := head
	for pre != nil {
		if pre.Val != cur.Val {
			return false
		}

		pre = pre.Next
		cur = cur.Next
	}

	return true
}

func TestParili(t *testing.T) {
	list := CreateListBySlice([]int{1, 2, 3, 3, 2, 1})

	res := isPalindrome(list)

	fmt.Println(res)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 错误演示，这里第一遍遍历的时候，修改了链表的位置，因此会出现问题，咱们应该使用一遍解决问题的方式
func oddEvenList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	dummyO, dummyE := &ListNode{}, &ListNode{}

	odd, even := dummyO, dummyE
	cur := head
	for cur != nil && cur.Next != nil {
		odd.Next = cur
		odd = odd.Next
		cur = cur.Next.Next
	}

	odd.Next = nil

	cur = head.Next
	fmt.Println("+++++++")
	PrintlnList(head)
	for cur != nil && cur.Next != nil {
		even.Next = cur
		even = even.Next
		cur = cur.Next.Next
		fmt.Println("ok")
	}

	PrintlnList(dummyO)
	PrintlnList(dummyE)

	even.Next = nil

	odd.Next = dummyE.Next

	return dummyO.Next
}

func TestOddEven(t *testing.T) {
	list := CreateListBySlice([]int{1, 2, 3, 4})
	res := oddEvenList(list)

	PrintlnList(res)
}

func oddEvenList(head *ListNode) *ListNode {
	oddDummy, evenDummy := &ListNode{}, &ListNode{}
	odd, even := oddDummy, evenDummy

	cnt := 1
	cur := head
	for cur != nil {
		if cnt%2 == 0 {
			even.Next = cur
			even = cur
		} else {
			odd.Next = cur
			odd = cur
		}

		cnt++
		cur = cur.Next
	}

	// 注意，需要处理某位节点
	even.Next = nil
	odd.Next = evenDummy.Next

	return oddDummy.Next
}

func Test_OddEvenList(t *testing.T) {
	list := CreateListBySlice([]int{1, 2, 3})
	PrintlnList(oddEvenList(list))
}
