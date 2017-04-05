package gofaker

import "strings"

type Lorem struct {
	faker *Faker
}

func (lorem *Lorem) Word() string {
	return lorem.faker.MustParse("lorem.words")
}

func (lorem *Lorem) Words() string {
	return lorem.WordsNum(3)
}

func (lorem *Lorem) WordsNum(numWords int) string {
	word_list := lorem.faker.data.Get("lorem.words")
	return strings.Join(lorem.faker.Sample(word_list, numWords), " ")
}

func (lorem *Lorem) WordsSupplemental(numWords int) string {
	word_list := lorem.faker.data.Get("lorem.words")
	word_list = append(word_list, lorem.faker.data.Get("lorem.supplemental")...)
	return strings.Join(lorem.faker.Sample(word_list, numWords), " ")
}

func (lorem *Lorem) Characters(num int) string {
	var result []rune
	for i := 0; i < num; i++ {
		if lorem.faker.random.Int() % 3 != 0 {
			result = append(result, lorem.faker.RandomLetter())
		} else {
			result = append(result, lorem.faker.RandomDigit())
		}
	}
	return string(result)
}

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

func (lorem *Lorem) Sentence() string {
	return lorem.SentenceNum(4, 6)
}

func (lorem *Lorem) SentencesNum(numSentences int) string {
	var sentences []string
	for i := 0; i < numSentences; i++ {
		sentences = append(sentences, lorem.SentenceNum(3, 6))
	}
	return strings.Join(sentences, " ")
}


func (lorem *Lorem) ParagraphNum(numSentences int, randomSentencesToAdd int) string {
	return lorem.SentencesNum(numSentences + lorem.faker.random.Intn(randomSentencesToAdd))
}

func (lorem *Lorem) Paragraph() string {
	return lorem.ParagraphNum(3, 6)
}

func (lorem *Lorem) Paragraphs(numParagraphs int) []string {
	var result []string
	for i := 0; i < numParagraphs; i++ {
		result = append(result, lorem.Paragraph())
	}
	return result
}
