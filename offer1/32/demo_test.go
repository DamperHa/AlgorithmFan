package _2

import (
	"fmt"
	"testing"
)

// 实现层次遍历算法

func TestBFS(t *testing.T) {
	// 通过数组构建一棵二叉树
	nums := []int{1, 2, 3, 4, 5}
	root := createBinaryTree(nums, 0, len(nums)-1)

	res := levelOrder(root)

	fmt.Println(res)
}

// 给定一个有序数组，构建一个平衡二叉树
func createBinaryTree(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2
	root := &TreeNode{
		val: nums[mid],
	}

	root.left = createBinaryTree(nums, left, mid-1)
	root.right = createBinaryTree(nums, mid+1, right)
	return root
}

type TreeNode struct {
	val   int       // 数据域
	left  *TreeNode // 指针域
	right *TreeNode
}

// levelOrder 层次遍历算法
// 所有数据都往queue里面放，只从queue中拿数据
// 可以将整个过程想象为一个放数据和拿数据的过程
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) != 0 {
		levelLen := len(queue)
		levelRes := []int{}
		for i := 0; i < levelLen; i++ {
			tmp := queue[0]
			queue = queue[1:]

			if tmp.left != nil {
				queue = append(queue, tmp.left)
			}

			if tmp.right != nil {
				queue = append(queue, tmp.right)
			}

			levelRes = append(levelRes, tmp.val)
		}

		res = append(res, levelRes)
	}

	return res
}
