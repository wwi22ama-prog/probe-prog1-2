package aufgabe5

import "fmt"

func ExampleDictionary_HasDuplicates() {
	d1 := Dictionary{
		[]Entry{
			{De: "Haus", En: "house"},
			{De: "Fahrrad", En: "bicycle"},
			{De: "Auto", En: "car"},
			{De: "Fahrrad", En: "bike"},
		},
	}

	d2 := Dictionary{
		[]Entry{
			{De: "Haus", En: "house"},
			{De: "Fahrrad", En: "bicycle"},
			{De: "Auto", En: "car"},
			{De: "Fahrrad", En: "bicycle"},
		},
	}

	fmt.Println(d1.HasDuplicates())
	fmt.Println(d2.HasDuplicates())

	// Output:
	// false
	// true
}
