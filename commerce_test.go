package gofaker

import "fmt"

func ExampleCommerce_Color() {
	faker, _ := NewFaker("en")
	faker.Reset()
	out := faker.Commerce.Color()
	fmt.Println(out)
	// Output: yellow
}

func ExampleCommerce_Department() {
	faker, _ := NewFaker("en")
	faker.Reset()
	out := faker.Commerce.Department()
	fmt.Println(out)
	// Output: Beauty
}

func ExampleCommerce_ProductName() {
	faker, _ := NewFaker("en")
	faker.Reset()
	out := faker.Commerce.ProductName()
	fmt.Println(out)
	// Output: Incredible Plastic Hat
}

func ExampleCommerce_Price() {
	faker, _ := NewFaker("en")
	faker.Reset()
	out := faker.Commerce.Price()
	fmt.Printf("%.2f\n", out)
	// Output: 37.30
}
