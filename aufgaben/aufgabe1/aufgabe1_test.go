package aufgabe1

import "fmt"

func ExampleDivisors_common_case() {
	fmt.Println(Divisors(10))
	fmt.Println(Divisors(100))
	fmt.Println(Divisors(55))
	fmt.Println(Divisors(11))

	// Output:
	// [1 2 5 10]
	// [1 2 4 5 10 20 25 50 100]
	// [1 5 11 55]
	// [1 11]
}

// Test von Divisors für den Fall n == 0.
func ExampleDivisors_zero_case() {
	// Das Ergebnis wird nicht ausgegeben, weil keine Anforderungen definiert sind.
	// Der Aufruf findet dennoch statt, um sehen zu können, ob die Funktion abstürzt.
	Divisors(0)

	// Output:
}
