package aufgabe2

import "fmt"

func ExampleArrayProducts() {
	fmt.Println(ArrayProducts([]int{1, 3, 5, 7}))
	fmt.Println(ArrayProducts([]int{1, 1, 2, 80}))
	fmt.Println(ArrayProducts([]int{7, 3, 1, 2}))
	fmt.Println(ArrayProducts([]int{3, 3, 0, 2}))

	// Output:
	// [1 3 15 105]
	// [1 1 2 160]
	// [7 21 21 42]
	// [3 9 0 0]
}
