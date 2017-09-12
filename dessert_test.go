package gofaker

import "fmt"

func ExampleDessert_Variety() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Dessert.Variety())
	// Output: Sweet Bread
}

func ExampleDessert_Topping() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Dessert.Topping())
	// Output: Bacon
}

func ExampleDessert_Flavor() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Dessert.Flavor())
	// Output: Banana
}
