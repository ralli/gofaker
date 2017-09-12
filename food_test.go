package gofaker

import "fmt"

func ExampleFood_Dish() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Food.Dish())
	// Output: Scotch eggs
}

func ExampleFood_Ingredient() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Food.Ingredient())
	// Output: Blue Cheese
}

func ExampleFood_Spice() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Food.Spice())
	// Output: Cumin Seed Black
}

func ExampleFood_Measurement() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Food.Measurement())
	// Output: 3 gallon
}

func ExampleFood_MetricMeasurement() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Food.MetricMeasurement())
	// Output: deciliter
}
