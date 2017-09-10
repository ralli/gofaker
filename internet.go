package gofaker

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var userSeparators = []string{"_", "."}
var reUserSplit = regexp.MustCompile(`\W+`)
var safeEmailSuffixes = []string{"org", "com", "net"}
var privateNetRegEx = []*regexp.Regexp{
	regexp.MustCompile(`^10\.`),
	regexp.MustCompile(`^127\.`),
	regexp.MustCompile(`^169\.254\.`),
	regexp.MustCompile(`^172\.(16|17|18|19|2\d|30|31)\.`),
	regexp.MustCompile(`^192\.168\.`),
}

// Internet provides functions to generate fake data for email addresses, user names, URLs and the like.
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
	internet.faker.shuffleStrings(a)
	return internet.UserWithName(strings.Join(a, " "))
}

var passwordChars = []rune("0123456789abcdefghjiklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!$%&/()=?-_.:,;#'+*<>")

// PasswordWithLength returns a password with a length ranging from minLength to maxLength
func (internet *Internet) PasswordWithLength(minLength int, maxLength int) string {
	length := internet.faker.random.Intn(maxLength-minLength) + minLength
	result := make([]rune, 0, length)
	for i := 0; i < length; i++ {
		idx := internet.faker.random.Intn(len(passwordChars))
		result = append(result, passwordChars[idx])
	}
	return string(result)
}

// Password returns a password
func (internet *Internet) Password() string {
	return internet.PasswordWithLength(6, 10)
}

// EmailWithUserName returns an email-Address for a given name
func (internet *Internet) EmailWithUsername(name string) string {
	return internet.UserWithName(name) + "@" + internet.DomainWord() + "." + internet.faker.MustParse("internet.domain_suffix")
}

// Email returns an email address
func (internet *Internet) Email() string {
	return internet.UserName() + "@" + internet.DomainWord() + "." + internet.faker.MustParse("internet.domain_suffix")
}

// FreeMailWithName generates an email address for a given user at a known free mail provider like gmail.
// The given user name will be normalized so that it becomes a valid email user name.
func (internet *Internet) FreeMailWithName(name string) string {
	return internet.UserWithName(name) + "@" + internet.faker.MustParse("internet.free_email")
}

// FreeMail generates an email address for a random user at a known free mail provider like gmail.
// The given user name will be normalized so that it becomes a valid email user name.
func (internet *Internet) FreeMail() string {
	return internet.UserName() + "@" + internet.faker.MustParse("internet.free_email")
}

// SafeMailWithName generates an email address for which it is safe to send test mail to (like example.com).
// The given user name will be normalized so that it becomes a valid email user name.
func (internet *Internet) SafeMailWithName(name string) string {
	return internet.UserWithName(name) + "@example." + internet.faker.randomValue(safeEmailSuffixes)
}

// SafeMailWithName generates an email for which it is safe to send test mail to (like example.com).
func (internet *Internet) SafeMail() string {
	return internet.UserName() + "@example." + internet.faker.randomValue(safeEmailSuffixes)
}

// DomainWord returns a valid domain part of an URL.
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

// MacAddress generates a MAC address.
func (internet *Internet) MacAddress() string {
	var parts []string
	for i := 0; i < 6; i++ {
		parts = append(parts, fmt.Sprintf("%02x", internet.faker.random.Intn(256)))
	}
	return strings.Join(parts, ":")
}

// IPv4Address generates an IP-Address.
func (internet *Internet) IPv4Address() string {
	var parts []string
	for i := 0; i < 4; i++ {
		parts = append(parts, fmt.Sprintf("%d", internet.faker.random.Intn(253)+2))
	}
	return strings.Join(parts, ".")
}

// PrivateIPv4Address generates an IP Address reserved for the local network like "192.168.x.y".
func (internet *Internet) PrivateIPv4Address() string {
	for {
		s := internet.IPv4Address()
		if isPrivateNet(s) {
			return s
		}
	}
}

func isPrivateNet(s string) bool {
	for _, r := range privateNetRegEx {
		if r.MatchString(s) {
			return true
		}
	}
	return false
}

// IPv6Address returns a valid IP V6 Address.
func (internet *Internet) IPv6Address() string {
	var parts []string

	for i := 0; i < 8; i++ {
		parts = append(parts, fmt.Sprintf("%x", internet.faker.random.Intn(65536)))
	}

	return strings.Join(parts, ":")
}
