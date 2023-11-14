package _53

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 给定一个数组，和下标n，返回下标对应数出现的次数
func GetNumberOfK(nums []int, k int) int {
	left := findLeft(nums, k)
	right := findRight(nums, k)

	res := right - left + 1

	return res
}

// 迭代找最左边
func findLeft(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == nums[k] {
			// 最左边数有什么特点呢？mid == - || nums[mid - 1] != nums[k]
			if mid == 0 || nums[mid-1] != nums[k] {
				return mid
			} else {
				right = mid - 1
			}
		} else if nums[mid] < nums[k] {
			left = mid + 1
		} else if nums[k] < nums[mid] {
			right = mid - 1
		}
	}

	return 0
}

// 迭代找最右边
func findRight(nums []int, k int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == nums[k] {
			// 最右边一个k有什么特点呢，mid = len(nums) - 1 || numss[mid + 1 ] != nums[k]
			if mid == len(nums)-1 || nums[mid+1] != nums[k] {
				return mid
			} else {
				left = mid + 1
			}

		} else if nums[mid] < nums[k] {
			left = mid + 1
		} else if nums[k] < nums[mid] {
			right = mid - 1
		}
	}

	return 0
}

// 找最右边那个数
func DFSRight(nums []int, k, left, right int) int {
	// 递归终止条件
	if left > right {
		return -1
	}

	// 当层处理逻辑
	mid := left + (right-left)/2
	if nums[mid] == nums[k] {
		if mid == len(nums)-1 || nums[mid] != nums[mid+1] {
			return mid
		} else {
			left = mid + 1
		}
	} else if nums[mid] < nums[k] {
		left = mid + 1
	} else if nums[k] < nums[mid] {
		right = mid - 1
	}

	// 进入下一层
	// 传递结果
	return DFSRight(nums, k, left, right)
}

func TestDFSRight(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 3, 4, 5}
	res := DFSRight(nums, 2, 0, len(nums)-1)

	assert.Equal(t, 5, res)
}

func Test53(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 3, 4, 5}

	// 中间数
	res1 := GetNumberOfK(nums, 2)
	assert.Equal(t, 4, res1)

	// 最左边
	res2 := GetNumberOfK(nums, 7)
	assert.Equal(t, 1, res2)

	// 最右边
	res3 := GetNumberOfK(nums, 0)
	assert.Equal(t, 1, res3)
}
