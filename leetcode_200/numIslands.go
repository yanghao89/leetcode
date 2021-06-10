package leetcode_200

func helper(grid [][]byte, H, L int) {
	hn := len(grid)
	ln := len(grid[0])
	if H < 0 || L < 0 || H >= hn || L >= ln || grid[H][L] == '0' {
		return
	}
	grid[H][L] = '0'
	helper(grid, H-1, L)
	helper(grid, H+1, L)
	helper(grid, H, L-1)
	helper(grid, H, L+1)
}

func NumIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	nr, nc, numIslands := len(grid), len(grid[0]), 0
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				numIslands++
				helper(grid, r, c)
			}
		}
	}
	return numIslands
}
