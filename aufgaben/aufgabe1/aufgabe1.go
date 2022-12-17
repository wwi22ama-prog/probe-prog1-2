/* AUFGABE: Listen
 * Erreichbare Punkte: 10
 */
package aufgabe1

// AUFGABENSTELLUNG:
// Erwartet eine Zahl n != 0 und liefert eine Liste mit allen Teilern von n.
// Für n == 0 ist das Verhalten nicht definiert, aber die Funktion darf nicht abstürzen.
func Divisors(n int) []int {
	result := []int{}
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			result = append(result, i)
		}
	}
	return result
}
