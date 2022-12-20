/* AUFGABE: Rekursion
 * Erreichbare Punkte: 15
 */
package aufgabe7

// AUFGABENSTELLUNG:
// Implementieren Sie die folgende Funktion.
// Zusatzanforderung: Die Funktion muss rekursiv bleiben.

// Berechnet den Rest bei der Division x / y.
// Soll sich genau so verhalten, wie der eingebaute Modulo-Operator (vgl. Tests).
func Mod(x, y int) int {
	// Sonderfälle:
	if x < 0 {
		return -Mod(-x, y)
	}
	if y < 0 {
		return Mod(x, -y)
	}

	// Ansatz:
	// Wir ziehen y von x ab, bis es nicht mehr geht.
	// Dann liefern wir den Rest.

	// Basisfall: "Es geht nicht mehr."
	// Genauer: Wenn y > x ist, dann können wir nichts mehr abziehen.
	//          Dann ist x selbst der Rest.
	if x < y {
		return x
	}
	return Mod(x-y, y)
}
