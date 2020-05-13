// Package tic_tac_toe which holds most of the codebase for this programd
package tic_tac_toe

type board [3][3]string

// newBoard is a constructor for an empty board
func newBoard() *board {

	var b board
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			b[i][j] = "_"
		}
	}
	return &b
}

// String returns the string representation of a board.
func (b *board) String() string {
	str := "\n    "
	for i := range b {
		for j := range b[i] {
			str += b[i][j] + " "
		}
		str += "\n    "
	}
	return str
}

// This method return empty cell's count of created grid
func (b *board) emptyCount() int {
	count := 0
	for i := range b {
		for j := range b[i] {
			if b[i][j] == "_" {
				count++
			}
		}
	}
	return count
}

// This method return whether match is won or nor
func (b *board) isWon() bool {
	if b.winner() != "" {
		return true
	}
	return false
}

// Winner returns the winner of the board.
func (b *board) winner() string {

	xStreak := [3]string{"X", "X", "X"}
	oStreak := [3]string{"O", "O", "O"}

	rows := b.rows()
	for i := range rows {
		if rows[i] == xStreak {
			return "X"
		}
		if rows[i] == oStreak {
			return "O"
		}
	}

	columns := b.columns()
	for i := range columns {
		if columns[i] == xStreak {
			return "X"
		}
		if columns[i] == oStreak {
			return "O"
		}
	}

	diagonals := b.diagonals()
	for i := range diagonals {
		if diagonals[i] == xStreak {
			return "X"
		}
		if diagonals[i] == oStreak {
			return "O"
		}
	}

	return ""
}

//initializes rows for given grid
func (b *board) rows() [3][3]string {
	return *b
}

//initializes columns for given grid
func (b *board) columns() [3][3]string {
	columns := [3][3]string{}
	for i := range b {
		for j := range b[i] {
			columns[j][i] = b[i][j]
		}
	}
	return columns
}

//initializes diagonals for given grid
func (b *board) diagonals() [2][3]string {
	diagonals := [2][3]string{}
	for i := range b {
		diagonals[0][i] = b[i][i]
	}

	for i := range b {
		diagonals[1][i] = b[i][2-i]
	}
	return diagonals
}
