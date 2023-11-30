package arr

import "math/rand"

func randomOne(nums []int) int {
	// 这里的n无法确定
	cnt := 0
	ans := 0
	for _, num := range nums {
		cnt++
		if rand.Intn(cnt) == 0 {
			ans = num
		}
	}

	return ans
}

func randomK(nums []int, k int) []int {
	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, nums[i])
	}

	cnt := k
	for i := k; i < len(nums); i++ {
		cnt++
		r := rand.Intn(cnt) // 【0, cnt-1】在这个区间取值
		if r < k {
			res[r] = nums[i]
		}
	}

	return res
}
