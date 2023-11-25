package _6

import (
	"fmt"
	"strconv"
	"testing"
)

func TestName(t *testing.T) {
	str := "226"
	fmt.Println(dfs(str, len(str)-1))
}

func dfs(str string, i int) int {
	if i == 0 {
		return 1
	}

	if i < 0 {
		return 0
	}

	// 当层的结果
	var res int

	res += dfs(str, i-1)

	item := str[i-1 : i+1]
	itemInt, _ := strconv.Atoi(item)
	if itemInt <= 25 {
		res += dfs(str, i-2)

	}

	return res
}

func TestSpace(t *testing.T) {
	i := ' '
	fmt.Println(i - '0') // -16
}

