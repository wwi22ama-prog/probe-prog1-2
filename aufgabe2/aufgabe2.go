/* AUFGABE: Listen
 * Erreichbare Punkte: 10
 */
package aufgabe2

// AUFGABENSTELLUNG:
// Erwartet eine int-Slice list.
// Liefert eine int-Slice, die an Stelle n das Produkt
// der Elemente aus list bis zu Stelle n enthält.
func ArrayProducts(list []int) []int {
	result := []int{}
	product := 1
	for _, el := range list {
		product *= el
		result = append(result, product)
	}
	return result
}
