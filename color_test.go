package gofaker

import "fmt"

func ExampleColor_Name() {
	faker, _ := NewFaker("en")
	faker.Reset()
	fmt.Println(faker.Color.Name())
	// Output: yellow
}
