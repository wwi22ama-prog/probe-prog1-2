package aufgabe8

type SudokuBoard [][]int

// Liefert ein leeres Spielfeld.
// Leere Zellen haben standardmäßig den Wert -1.
func EmptySudokuBoard() SudokuBoard {
	return SudokuBoard{
		{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1},
	}
}

// Prüft, ob eine gegebene Zelle leer ist.
func (board SudokuBoard) CellEmpty(row, column int) bool {
	return board[row][column] == -1
}

// Schreibt einen Wert in eine gegebene Zelle, falls dort noch keiner gesetzt ist.
func (board *SudokuBoard) SetValue(row, column, value int) {
	if board.CellEmpty(row, column) {
		(*board)[row][column] = value
	}
}
