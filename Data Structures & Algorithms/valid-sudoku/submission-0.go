func isValidSudoku(board [][]byte) bool {
	//check rows
	hash_set := make(map[byte]struct{})
	for i:= 0; i<9; i++ {
		for j:=0; j<9; j++ {
			if board[i][j] == 46 {
				continue
			}

			if _, exists := hash_set[board[i][j]]; exists {
				return false
			}
			hash_set[board[i][j]] = struct{}{}
		}
		clear(hash_set)
	}

	fmt.Println("passed row check")
	//check cols
	for i:= 0; i<9; i++ {
		for j:=0; j<9; j++ {
			if board[j][i] == 46 {
				continue
			}

			if _, exists := hash_set[board[j][i]]; exists {
				return false
			}
			hash_set[board[j][i]] = struct{}{}
		}
		clear(hash_set)
	}

	fmt.Println("passed cols check")
	//check box
	for x:=0 ; x<9; x+=3 {
		for y:=0; y<9; y+=3 {
			for i:=0; i<3; i++ {
				for j:=0; j<3; j++{
					if board[x+i][y+j] == 46 {
						continue
					}

					if _, exists := hash_set[board[x+i][y+j]]; exists {
						return false
					}
					hash_set[board[x+i][y+j]] = struct{}{}
				}
			}
			clear(hash_set)
		}
	}

	return true
}
