package aufgabe6

import "fmt"

func ExampleDictionary_AllTranslations() {
	d1 := Dictionary{
		[]Entry{
			{De: "Haus", En: "house"},
			{De: "Fahrrad", En: "bicycle"},
			{De: "Auto", En: "car"},
			{De: "Fahrrad", En: "bike"},
		},
	}

	fmt.Println(d1.AllTranslations("Haus"))
	fmt.Println(d1.AllTranslations("Fahrrad"))
	fmt.Println(d1.AllTranslations("Kuh"))
	// Output:
	// [house]
	// [bicycle bike]
	// []
}
