package gofaker

import "fmt"

func ExampleRomanize() {
	in := "abcäöüÄÖÜß"
	out := Romanize(in)
	fmt.Println(out)
	// Output: abcaeoeueAeOeUess
}
