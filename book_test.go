package gofaker

import "fmt"

func ExampleBook_Author() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Book.Author())
	// Output: Camilla Swift
}

func ExampleBook_Genre() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Book.Genre())
	// Output: Fanfiction
}

func ExampleBook_Publisher() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Book.Publisher())
	// Output: Pen and Sword Books
}

func ExampleBook_Title() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Book.Title())
	// Output: Alone on a Wide, Wide Sea
}
