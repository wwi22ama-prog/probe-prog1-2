/* AUFGABE: Rekursion
 * Erreichbare Punkte: 15
 */
package aufgabe7

// AUFGABENSTELLUNG:
// Reparieren Sie die folgende Funktion.
// Es müssen alle Tests grün werden.
// Zusatzanforderung: Die Funktion muss rekursiv bleiben.

// Berechnet den Rest bei der Division x / y.
// Soll sich genau so verhalten, wie der eingebaute Modulo-Operator (vgl. Tests).
func Mod(x, y int) int {
	if x < 0 {
		return -Mod(-x, y)
	}
	if y < 0 {
		return Mod(x, -y)
	}
	if x < y {
		return x
	}
	return Mod(x-y, y)
}
