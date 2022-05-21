package minesweeper

// Annotate returns an annotated board

func Annotate(board []string) []string {
    new_board := make([][]byte,0)
    m := len(board)
	if m == 0 {
		return board
	}
	n := len(board[0])
    for _, str := range board {
        new_board = append(new_board, []byte(str))
    }

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if new_board[row][col] == byte('*') {
				continue
			}
			value := countMineNeighbours(new_board, row, col)
			if value > 0 {
				new_board[row][col] = byte('0' + value)
			}
		}
	}

	result := make([]string,0)
	for _, row := range new_board {
		s := ""
		for _, char := range row {
			s += string(char)
		}
		result = append(result, s)
	}
	return result
} 

func countMineNeighbours(board [][]byte, row, col int) int {
    rows := len(board)
    cols := len(board[0])
    neighbours := 0
    directions := [][]int{
        {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}  	}
    for _, dirs := range directions {
        x := dirs[0]
		y := dirs[1]
		new_x := row + x
		new_y := col + y
        if (0 <= new_x && new_x < rows) &&  (0 <= new_y && new_y < cols) && board[new_x][new_y] == byte('*'){
            neighbours++
        }
    }
    return neighbours
}
