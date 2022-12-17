/* AUFGABE: Structs
 * Erreichbare Punkte: 15
 */
package aufgabe5

// AUFGABENSTELLUNG:
// Reparieren Sie die folgende Methode.
// Es müssen alle Tests grün werden.

// Prüft, ob dict Duplikate enthält.
func (dict Dictionary) HasDuplicates() bool {
	for i, entry1 := range dict.Entries {
		for _, entry2 := range dict.Entries[i+1:] {
			if entry1 == entry2 {
				return true
			}
		}
	}
	return false
}
