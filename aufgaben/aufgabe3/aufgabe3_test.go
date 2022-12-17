package aufgabe3

import "fmt"

func ExampleIsPowerOf2() {

	fmt.Println(IsPowerOf2(2))
	fmt.Println(IsPowerOf2(5))
	fmt.Println(IsPowerOf2(0))
	fmt.Println(IsPowerOf2(64))
	fmt.Println(IsPowerOf2(-2))
	fmt.Println(IsPowerOf2(16384))

	// Output:
	// true
	// false
	// false
	// true
	// false
	// true
}
