package gofaker

import "strings"

// Lorem provides functions for fake data related to "Lorem Ipsum" sentences.
type Lorem struct {
	faker *Faker
}

// Word generates a single Lorem Ipsum word.
func (lorem *Lorem) Word() string {
	return lorem.faker.MustParse("lorem.words")
}

// Words generates three Lorem words separated by whitespace.
func (lorem *Lorem) Words() string {
	return lorem.WordsNum(3)
}

// Words generates a given number of lorem words separated by spaces.
func (lorem *Lorem) WordsNum(numWords int) string {
	word_list := lorem.faker.data.Get("lorem.words")
	return strings.Join(lorem.faker.Sample(word_list, numWords), " ")
}

// WordsSupplemental generates a given number of lorem words using an extendes set of words separated by spaces.
func (lorem *Lorem) WordsSupplemental(numWords int) string {
	word_list := lorem.faker.data.Get("lorem.words")
	word_list = append(word_list, lorem.faker.data.Get("lorem.supplemental")...)
	return strings.Join(lorem.faker.Sample(word_list, numWords), " ")
}

// Characters genrates a sequence of alphanumeric characters.
func (lorem *Lorem) Characters(num int) string {
	var result []rune
	for i := 0; i < num; i++ {
		if lorem.faker.random.Int()%3 != 0 {
			result = append(result, lorem.faker.RandomLetter())
		} else {
			result = append(result, lorem.faker.RandomDigit())
		}
	}
	return string(result)
}

// SentenceNum generates a sentence with a minimum number of words. A sentence starts with an uppercase first
// word and ends with a period. The number of words may vary by a random number of additional words.
func (lorem *Lorem) SentenceNum(numWords int, randomWordsToAdd int) string {
	var words []string

	for i := 0; i < numWords; i++ {
		words = append(words, lorem.Word())
	}

	n := lorem.faker.random.Intn(randomWordsToAdd)
	for i := 0; i < n; i++ {
		words = append(words, lorem.Word())
	}

	words[0] = strings.Title(words[0])
	return strings.Join(words, " ") + "."
}

// Returns a sentence. A sentence starts with an uppercase first
// word and ends with a period.
func (lorem *Lorem) Sentence() string {
	return lorem.SentenceNum(4, 6)
}

// SentencesNum returns a given number of sentences.
func (lorem *Lorem) SentencesNum(numSentences int) string {
	var sentences []string
	for i := 0; i < numSentences; i++ {
		sentences = append(sentences, lorem.SentenceNum(3, 6))
	}
	return strings.Join(sentences, " ")
}

// ParagraphNum returns a paragraph with a minimum number of sentences. There may be a random number of sentences
// which will be added to the paragraph.
func (lorem *Lorem) ParagraphNum(numSentences int, randomSentencesToAdd int) string {
	return lorem.SentencesNum(numSentences + lorem.faker.random.Intn(randomSentencesToAdd))
}

// Paragraph returns a paragraph.
func (lorem *Lorem) Paragraph() string {
	return lorem.ParagraphNum(3, 6)
}

// Paragraphs returns a list of paragraphs.
func (lorem *Lorem) Paragraphs(numParagraphs int) []string {
	var result []string
	for i := 0; i < numParagraphs; i++ {
		result = append(result, lorem.Paragraph())
	}
	return result
}
