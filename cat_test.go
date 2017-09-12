package gofaker

import "fmt"

func ExampleCat_Name() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Cat.Name())
	// Output: Coco
}

func ExampleCat_Breed() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Cat.Breed())
	// Output: British Semipi-longhair
}

func ExampleCat_Registry() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Cat.Registry())
	// Output: American Cat Fanciers Association
}
