package aufgabe7

import "fmt"

func ExampleMod_print() {
	// Anmerkung: Die Funktion soll sich in allen Belangen so verhalten,
	// wie der eingebaute Modulo-Operator.
	// Dies wird hier im Test zusätzlich zu geprüft.

	fmt.Println(Mod(3, 2))
	fmt.Println(Mod(-3, 2))
	fmt.Println(Mod(3, -2))
	fmt.Println(Mod(-3, -2))

	// Output:
	// 1
	// -1
	// 1
	// -1
}

// Die Funktion soll sich in allen Belangen so verhalten,
// wie der eingebaute Modulo-Operator.
// Dies wird in diesem Test geprüft.
func ExampleMod_operator() {

	fmt.Println(Mod(3, 1) == 3%1)         // == 0
	fmt.Println(Mod(3, 2) == 3%2)         // == 1
	fmt.Println(Mod(3, 3) == 3%3)         // == 0
	fmt.Println(Mod(3, 4) == 3%4)         // == 3
	fmt.Println(Mod(-3, 2) == (-3)%2)     // == -1
	fmt.Println(Mod(3, -2) == 3%(-2))     // == 1
	fmt.Println(Mod(-3, -2) == (-3)%(-2)) // == -1

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
	// true
}
