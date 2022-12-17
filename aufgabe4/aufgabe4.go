/* AUFGABE: Listen
 * Erreichbare Punkte: 10
 */
package aufgabe4

// AUFGABENSTELLUNG:
// Liefert eine Liste mit allen Werten aus list, außer x.
func FilterX(list []int, x int) []int {
	result := []int{}
	for _, el := range list {
		if el != x {
			result = append(result, el)
		}
	}
	return result
}
