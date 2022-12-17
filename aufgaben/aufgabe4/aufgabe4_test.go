package aufgabe4

import "fmt"

func ExampleFilterX() {
	l1 := []int{1, 4, 17, 2, 1, 5, 10, 5, 2}
	fmt.Println(FilterX(l1, 1))
	fmt.Println(FilterX(l1, 4))
	fmt.Println(FilterX(l1, 5))
	fmt.Println(FilterX(l1, 42))

	// Output:
	// [4 17 2 5 10 5 2]
	// [1 17 2 1 5 10 5 2]
	// [1 4 17 2 1 10 2]
	// [1 4 17 2 1 5 10 5 2]
}
