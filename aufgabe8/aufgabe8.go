/* AUFGABE: Sudoku
 * Erreichbare Punkte: 30
 */
package aufgabe8

// VORGABE:
// In der Datei sudokuboard.go ist ein Datentyp für ein Sudoku-Spielfeld vorgegeben:

// AUFGABENSTELLUNG:
// Erwartet Koordinaten sowie eine Zahl n.
// überprüft, ob es erlaubt ist, nach den Sudoku-Regeln n auf das angegebene Feld zu
// setzen.
func (board SudokuBoard) MoveAllowed(row, column, n int) bool {

	// Ungültige Eingaben und besetzte Zellen prüfen.
	if row < 0 || column < 0 || n < 0 || row > 8 || column > 8 || n > 8 || !board.CellEmpty(row, column) {
		return false
	}

	// Zeile prüfen.
	for _, value := range board[row] {
		if value == n {
			return false
		}
	}

	// Spalte prüfen.
	for r := range board {
		if board[r][column] == n {
			return false
		}
	}

	// Startpositionen des verbotenen Quadrats innerhalb des gesamten Felds bestimmen.
	startRow := row / 3 * 3
	startCol := column / 3 * 3

	// Verbotenes Quadrat überprüfen.
	for r := startRow; r < startRow+3; r++ {
		for c := startCol; c < startCol+3; c++ {
			if board[r][c] == n {
				return false
			}
		}
	}
	return true
}
