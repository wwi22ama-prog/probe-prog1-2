/* AUFGABE: Rekursion
 * Erreichbare Punkte: 10
 */
package aufgabe3

// AUFGABENSTELLUNG:
// Liefert true, falls x eine Zweierpotenz ist.
// Zusatzanforderung: Die Funktion muss rekursiv sein.
func IsPowerOf2(x int) bool {
	if x < 1 {
		return false
	}
	if x == 1 {
		return true
	}
	return x%2 == 0 && IsPowerOf2(x/2)
}
