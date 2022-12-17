package aufgabe8

import "fmt"

func ExampleSudokuBoard_MoveAllowed_row() {
	b1 := EmptySudokuBoard()

	b1.SetValue(3, 4, 5)

	// Zeile 3 ist belegt, es darf also keine 5 mehr in diese Zeile gesetzt werden.
	fmt.Println(b1.MoveAllowed(3, 2, 5))
	fmt.Println(b1.MoveAllowed(3, 8, 5))

	// Output:
	// false
	// false
}

func ExampleSudokuBoard_MoveAllowed_column() {
	b1 := EmptySudokuBoard()

	b1.SetValue(3, 4, 5)

	// Spalte 4 ist belegt, es darf also keine 5 mehr in diese Spalte gesetzt werden.
	fmt.Println(b1.MoveAllowed(6, 4, 5))
	fmt.Println(b1.MoveAllowed(2, 4, 5))

	// Output:
	// false
	// false
}

func ExampleSudokuBoard_MoveAllowed_forbiddenSquare() {
	b1 := EmptySudokuBoard()

	b1.SetValue(3, 4, 5)

	// Das Mittenquadrat ist belegt, es darf also keine 5 mehr dort gesetzt werden.
	fmt.Println(b1.MoveAllowed(3, 3, 5))
	fmt.Println(b1.MoveAllowed(3, 4, 5))
	fmt.Println(b1.MoveAllowed(3, 5, 5))
	fmt.Println(b1.MoveAllowed(4, 3, 5))
	fmt.Println(b1.MoveAllowed(4, 4, 5))
	fmt.Println(b1.MoveAllowed(4, 5, 5))
	fmt.Println(b1.MoveAllowed(5, 3, 5))
	fmt.Println(b1.MoveAllowed(5, 4, 5))
	fmt.Println(b1.MoveAllowed(5, 5, 5))

	// Stichprobenartige Prüfung erlaubter Positionen.

	// Es dürfen generell nur die Zahlen 0-8 gesetzt werden.

	// Output:
	// false
	// false
	// false
	// false
	// false
	// false
	// false
	// false
	// false
}

func ExampleSudokuBoard_MoveAllowed_allowedPosition() {
	b1 := EmptySudokuBoard()

	// Stichprobenartige Prüfung erlaubter Positionen.
	fmt.Println(b1.MoveAllowed(3, 3, 5))
	fmt.Println(b1.MoveAllowed(3, 7, 2))

	// Output:
	// true
	// true
}

func ExampleSudokuBoard_MoveAllowed_invalidMoves() {
	b1 := EmptySudokuBoard()

	// Es dürfen generell nur die Zahlen 0-8 gesetzt werden.
	// Es dürfen generell auch nur die Zeilen und Spalten 0-8 besetzt werden.
	fmt.Println(b1.MoveAllowed(3, 3, 15))
	fmt.Println(b1.MoveAllowed(3, 17, 2))
	fmt.Println(b1.MoveAllowed(13, 7, 2))

	// Output:
	// false
	// false
	// false
}
