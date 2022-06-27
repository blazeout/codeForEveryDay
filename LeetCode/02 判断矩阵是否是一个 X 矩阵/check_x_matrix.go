package main

/**
https://leetcode.cn/contest/weekly-contest-299/problems/check-if-matrix-is-x-matrix/
*/

func checkXMatrix(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == j && grid[i][j] == 0 {
				return false
			}
			if i == len(grid)-j-1 && grid[i][j] == 0 {
				return false
			}
			if i != j && i != len(grid)-1-j && grid[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	checkXMatrix([][]int{{2, 0, 0, 1}, {0, 3, 1, 0}, {0, 5, 2, 0}, {4, 0, 0, 2}})
}
