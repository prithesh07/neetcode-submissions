func searchMatrix(matrix [][]int, target int) bool {
	//O(log (m*n)) is O(log m) + O(log n) => one search along last column to find the matching row and then one search along that row to find


	n := len(matrix)
	m := len(matrix[0])
	// search along column 

	left := 0
	right := n - 1
	ans := 0

	for left <= right {
		mid := left + (right - left) / 2

		if matrix[mid][m-1] == target {
			return true
		} 
		
		if (matrix[mid][m-1] < target) {
			left = mid + 1
		} else {
			ans = mid
			right = mid - 1
		}
	}

	//search along the matching row
	left = 0
	right = m - 1

	for left <= right {
		mid := left + (right - left) / 2

		if matrix[ans][mid] == target {
			return true
		} 
		
		if (matrix[ans][mid] < target) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}


	return false
}
