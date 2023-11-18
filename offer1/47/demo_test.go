package _7

import (
	"testing"
)

func TestDemo(t *testing.T) {
	frame := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	println(jewelleryValue(frame))
}

func jewelleryValue(frame [][]int) int {

	m, n := len(frame), len(frame[0])
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 处理第一行的场景
	dp[0][0] = frame[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + frame[i][0]
	}

	// 处理第一列的场景
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + frame[0][j]
	}

	// 处理中间场景
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + frame[i][j]
		}
	}

	return dp[m-1][n-1]
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func ConstructArr(arr [][]int) [][]int {
	res := make([][]int, len(arr))

	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(arr[i]))
	}

	return res
}