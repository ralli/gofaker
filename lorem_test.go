package gofaker

import (
	"fmt"
)

func ExampleLorem_Word() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Lorem.Word())
	// Output: at
}

func ExampleLorem_Words() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Lorem.Words())
	// Output: at illum ut
}

func ExampleLorem_WordsNum() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Lorem.WordsNum(10))
	// Output: at illum ut est sit soluta nulla numquam nobis sunt
}

func ExampleLorem_WordsSupplemental() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Lorem.WordsSupplemental(10))
	// Output: cunabula color crepusculum tersus sunt decretum conor decumbo coadunatio sordeo
}

func ExampleLorem_Characters() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Lorem.Characters(10))
	// Output: R0T6Z78vn4
}

func ExampleLorem_Sentence() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Lorem.Sentence())
	// Output: At illum ut est soluta.
}

func ExampleLorem_SentencesNum() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Lorem.SentencesNum(3))
	// Output: At illum ut. Sit soluta nulla nobis sunt. Quaerat ea dolores deleniti culpa numquam ut.
}

func ExampleLorem_Paragraphs() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Lorem.Paragraphs(2))
	// Output: [Illum ut est soluta. Nulla numquam nobis quaerat. Ea dolores facere culpa numquam. Ut distinctio maxime est qui corporis sunt. Officia odit et odit. Molestias voluptas porro. Magnam ipsa corporis rem. Non hic esse quisquam hic. Earum molestias iste porro et blanditiis. Iste eum repellendus. Qui eius suscipit quia quo et nesciunt. Quod fuga ut pariatur libero sequi rerum. Omnis soluta facilis possimus et voluptas eaque possimus. Harum voluptatibus aperiam qui autem quam veniam. Ea voluptas facilis autem illum esse.]
}
