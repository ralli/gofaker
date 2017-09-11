package gofaker

import "fmt"

func ExampleApp_Name() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.App.Name())
	// Output: Bamity
}

func ExampleApp_Version() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.App.Version())
	// Output: 0.7.8
}

func ExampleApp_Author() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.App.Author())
	// Output: Crona, Swift and Zboncak

}
