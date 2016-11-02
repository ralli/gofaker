package gofaker

import (
	"bytes"
	"math/rand"
	"regexp"
	"strings"
	"unicode"
)

var userSeparators = []string{"_", "."}
var reUserSplit = regexp.MustCompile(`\W+`)

type Internet struct {
	faker *Faker
}

// UserWithName transforms a users name into something that can be used as a username (login).
func (internet *Internet) UserWithName(specifier string) string {
	in := strings.ToLower(Romanize(specifier))
	parts := reUserSplit.Split(in, -1)
	sep := userSeparators[internet.faker.random.Intn(len(userSeparators))]
	return strings.Join(parts, sep)
}

// UserName returns a random username.
func (internet *Internet) UserName() string {
	f, err := NewFaker("en")
	if err != nil {
		panic(err)
	}
	a := []string{f.Name.FirstName(), f.Name.LastName()}
	internet.faker.ShuffleStrings(a)
	return internet.UserWithName(strings.Join(a, " "))
}

var passwordChars = []rune("0123456789abcdefghjiklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!$%&/()=?-_.:,;#'+*<>")

func (internet *Internet) PasswordWithLength(minLength int, maxLength int) string {
	length := internet.faker.random.Intn(maxLength-minLength) + minLength
	result := make([]rune, 0, length)
	for i := 0; i < length; i++ {
		idx := rand.Intn(len(passwordChars))
		result = append(result, passwordChars[idx])
	}
	return string(result)
}

func (internet *Internet) Password() string {
	return internet.PasswordWithLength(6, 10)
}

func (internet *Internet) EmailWithUsername(name string) string {
	return internet.UserWithName(name) + "@" + internet.DomainWord() + internet.faker.MustParse("internet.domain_suffix")
}

func (internet *Internet) Email() string {
	return internet.EmailWithUsername(internet.faker.Name.FirstName() + " " + internet.faker.Name.LastName())
}

func (internet *Internet) FreeMailWithName(name string) string {
	return internet.UserWithName(name) + "@" + internet.faker.MustParse("internet.free_email")
}

func (internet *Internet) DomainWord() string {
	return strings.ToLower(internet.cleanupDomainName(internet.faker.Company.Name()))
}

func (internet *Internet) cleanupDomainName(in string) string {
	s := Romanize(in)
	var out bytes.Buffer
	first := true
	r := '_'
	if internet.faker.random.Intn(2) == 0 {
		r = '-'
	}
	for _, c := range s {
		if first {
			if unicode.IsLetter(c) {
				first = false
				out.WriteRune(c)
			}
		} else {
			if unicode.IsLetter(c) || unicode.IsDigit(c) {
				out.WriteRune(c)
			} else {
				out.WriteRune(r)
				first = true
			}
		}
	}
	return string(out.Bytes())
}
