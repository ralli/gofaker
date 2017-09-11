package gofaker

// Book provides functions related to books.
type Book struct {
	faker *Faker
}

// Title generates the books title.
func (b *Book) Title() string {
	return b.faker.MustParse("book.title")
}

// Author generates the books authors name.
func (b *Book) Author() string {
	return b.faker.MustParse("book.author")
}

// Publisher generates the books publishers name.
func (b *Book) Publisher() string {
	return b.faker.MustParse("book.publisher")
}

// Genre generates the books genre.
func (b *Book) Genre() string {
	return b.faker.MustParse("book.genre")
}
