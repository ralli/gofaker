package gofaker

import (
	"regexp"
	"strings"
)

var userSeparators = []string{"_", "."}
var reUserSplit = regexp.MustCompile(`\W+`)

type Internet struct {
	faker *Faker
}

func (internet *Internet) UserWithName(specifier string) string {
	in := strings.ToLower(Romanize(specifier))
	parts := reUserSplit.Split(in, -1)
	sep := userSeparators[internet.faker.random.Intn(len(userSeparators))]
	return strings.Join(parts, sep)
}

func (internet *Internet) UserName() string {
	f, err := NewFaker("en")
	if err != nil {
		panic(err)
	}
	a := []string{f.Name.FirstName(), f.Name.LastName()}
	internet.faker.ShuffleStrings(a)
	return internet.UserWithName(strings.Join(a, " "))
}

func (faker *Faker) ShuffleStrings(slice []string) {
	for i := range slice {
		j := faker.random.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
