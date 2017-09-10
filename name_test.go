package gofaker

import "fmt"

func ExampleName_FirstName() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Name.FirstName())
	// Output: Tommie
}

func ExampleName_LastName() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Name.LastName())
	// Output: Willms
}

func ExampleName_Name() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Name.Name())
	// Output: Evie Crona
}

func ExampleName_Title() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Name.Title())
	// Output: Global Tactics Producer
}
