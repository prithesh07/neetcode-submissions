func bfs(grid *[][]byte, visited *[][]bool, i int, j int, n int, m int) {
	if i < 0 || j < 0 || i >= n || j >= m || (*grid)[i][j] == '0' || (*visited)[i][j] {
		return
	}

	(*visited)[i][j] = true
	bfs(grid, visited, i+1, j, n, m)
	bfs(grid, visited, i, j+1, n, m)
	bfs(grid, visited, i-1, j, n, m)
	bfs(grid, visited, i, j-1, n, m)
}


func numIslands(grid [][]byte) int {
	n := len(grid)
	m := len(grid[0])

	visited := make([][]bool, n)

    for i := 0; i < n; i++ {
        visited[i] = make([]bool, m)
    }

	res := 0

    for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				res++
				bfs(&grid, &visited, i, j, n, m)
			}
		}
	}

	return res
}
