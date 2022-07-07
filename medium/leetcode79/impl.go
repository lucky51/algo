package leetcode79

import "fmt"

// exist 单词搜索
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	wl := len(word)
	var visited [][]bool
	visited = make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var check func(int, int, int) bool
	check = func(i, j, k int) bool {
		fmt.Println(i, j, k, string(word[k]))
		if word[k] != board[i][j] {
			return false
		}
		if k == wl-1 {
			return true
		}
		visited[i][j] = true
		defer func() { visited[i][j] = false }()
		if i-1 >= 0 && !visited[i-1][j] && check(i-1, j, k+1) {
			return true
		}
		if i+1 < m && !visited[i+1][j] && check(i+1, j, k+1) {
			return true
		}
		if j+1 < n && !visited[i][j+1] && check(i, j+1, k+1) {
			return true
		}
		if j-1 >= 0 && !visited[i][j-1] && check(i, j-1, k+1) {
			return true
		}
		//return check(i-1, j, k+1) || check(i+1, j, k+1) || check(i, j+1, k+1) || check(i, j-1, k+1)
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if check(i, j, 0) {
				return true
			}
		}

	}
	return false
}

// exist2 单词搜索
func exist2(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	wl := len(word)
	var visited [][]int
	visited = make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
	}
	var check func(int, int, int) bool
	check = func(i, j, k int) bool {
		// -。- 这东西会导致超时
		//	fmt.Println(i, j, k, string(word[k]))
		if word[k] != board[i][j] {
			return false
		}
		if k == wl-1 {
			return true
		}
		visited[i][j] = 1
		defer func() { visited[i][j] = 0 }()
		if i-1 >= 0 && visited[i-1][j] == 0 && check(i-1, j, k+1) {
			return true
		}
		if i+1 < m && visited[i+1][j] == 0 && check(i+1, j, k+1) {
			return true
		}
		if j+1 < n && visited[i][j+1] == 0 && check(i, j+1, k+1) {
			return true
		}
		if j-1 >= 0 && visited[i][j-1] == 0 && check(i, j-1, k+1) {
			return true
		}
		//return check(i-1, j, k+1) || check(i+1, j, k+1) || check(i, j+1, k+1) || check(i, j-1, k+1)
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word[0] == board[i][j] {
				if check(i, j, 0) {
					return true
				}
			}
		}

	}
	return false
}

// ==========================下边是官方解法================================================

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func exist1(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	vis := make([][]bool, h)
	for i := range vis {
		vis[i] = make([]bool, w)
	}
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] { // 剪枝：当前字符不匹配
			return false
		}
		if k == len(word)-1 { // 单词存在于网格中
			return true
		}
		vis[i][j] = true
		defer func() { vis[i][j] = false }() // 回溯时还原已访问的单元格
		for _, dir := range directions {
			if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < h && 0 <= newJ && newJ < w && !vis[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}
