package gofaker

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"
)

// Faker does something
type Faker struct {
	random     *rand.Rand
	data       fakerData
	Name       *Name
	Address    *Address
	App        *App
	Bank       *Bank
	Book       *Book
	CreditCard *CreditCard
	Phone      *Phone
	Internet   *Internet
	Color      *Color
	Company    *Company
	Compass    *Compass
	Code       *Code
	Commerce   *Commerce
	Lorem      *Lorem
}

var digits = []rune("0123456789")
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (faker *Faker) randomValue(a []string) string {
	length := len(a)
	if length == 0 {
		return "N/A"
	}
	return a[faker.random.Intn(length)]
}

func extractSubKeys(expression string) []string {
	const pattern = `#\{([^}]+)\}`
	re := regexp.MustCompile(pattern)
	a := re.FindAllStringSubmatch(expression, -1)
	var result []string
	for _, v := range a {
		result = append(result, v[1])
	}
	return result
}

// ToSnake converts the given string to snake case following the Golang format:
// acronyms are converted to lower-case and preceded by an underscore.
func ToSnake(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}

func getBaseKeyPrefix(k string) string {
	a := strings.Split(k, ".")
	if len(a) == 0 {
		return ""
	}
	return strings.Join(a[:len(a)-1], ".")
}

func (faker *Faker) subkeyValue(baseKey string, k string) (string, error) {
	// key contains a period => expand key directly
	if strings.Contains(k, ".") {
		kv := ToSnake(k)
		v, err := faker.Parse(kv)
		if err != nil {
			return "", err
		}
		return v, nil
	}
	// k does not contain a period => add the prefix of
	basePrefix := getBaseKeyPrefix(baseKey)
	kv := ToSnake(basePrefix + "." + k)
	v, err := faker.Parse(kv)
	if err != nil {
		kv = ToSnake(k)
		v, err = faker.Parse(kv)
		if err != nil {
			return "", err
		}
	}
	return v, nil
}

// Parse gets fake data given a key and expands variable values recursively.
func (faker *Faker) Parse(key string) (string, error) {
	v, err := faker.Fetch(key)
	if err != nil {
		return "", err
	}
	subKeys := extractSubKeys(v)
	for _, k := range subKeys {
		sv, err := faker.subkeyValue(key, k)
		if err != nil {
			return "", err
		}
		v = strings.Replace(v, "#{"+k+"}", sv, 1)
	}
	return v, nil
}

// MustParse returns the value of Parse and panics in case of an error
func (faker *Faker) MustParse(key string) string {
	v, err := faker.Parse(key)
	if err != nil {
		panic(err)
	}
	return v
}

// Fetch gets a fake data value without further expansion. That means, the result may
// contain unexpanded variables like '#{last_name}'
func (faker *Faker) Fetch(key string) (string, error) {
	v := faker.data.Get(key)
	if v == nil {
		return "", fmt.Errorf("%s not found", key)
	}
	return faker.randomValue(v), nil
}

// MustFetch gets a fake data value without further expansion. That means, the result may
// contain unexpanded variables like '#{last_name}'. Panics on error.
func (faker *Faker) MustFetch(key string) string {
	v, err := faker.Fetch(key)
	if err != nil {
		panic(err)
	}
	return v
}

// FetchOrEmpty gets a fake data value without further expansion. That means, the result may
// contain unexpanded variables like '#{last_name}'. Returns an empty string on error.
func (faker *Faker) FetchOrEmpty(key string) string {
	v, err := faker.Fetch(key)
	if err != nil {
		return ""
	}
	return v
}

func (faker *Faker) FetchList(key string) ([][]string, error) {
	v := faker.data.GetList(key)
	if v == nil {
		return nil, fmt.Errorf("%s not found", key)
	}
	return v, nil
}

func (faker *Faker) MustFetchList(key string) [][]string {
	v, err := faker.FetchList(key)
	if err != nil {
		panic(err)
	}
	return v
}

// NewFaker creates a new Faker instance for a given locale
func NewFaker(locale string) (*Faker, error) {
	data, err := newData(locale)
	if err != nil {
		return nil, err
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	faker := &Faker{random: random, data: data}
	faker.Name = &Name{faker}
	faker.Address = &Address{faker}
	faker.App = &App{faker}
	faker.Bank = &Bank{faker}
	faker.Book = &Book{faker}
	faker.CreditCard = &CreditCard{faker}
	faker.Phone = &Phone{faker}
	faker.Internet = &Internet{faker}
	faker.Company = &Company{faker}
	faker.Code = &Code{faker}
	faker.Color = &Color{faker}
	faker.Compass = &Compass{faker}
	faker.Commerce = &Commerce{faker}
	faker.Lorem = &Lorem{faker}
	return faker, nil
}

// Reset resets the fakers random number generator so that a repeatable sequence of values is returned.
// Useful for testing.
func (faker *Faker) Reset() {
	faker.random.Seed(42)
}

// RandomDigit returns a random digit.
func (faker *Faker) RandomDigit() rune {
	return digits[faker.random.Intn(10)]
}

// Numerify replaces all hash signs within a string with random digits.
func (faker *Faker) Numerify(s string) string {
	var out []rune
	for _, c := range s {
		if c == '#' {
			out = append(out, faker.RandomDigit())
		} else {
			out = append(out, c)
		}
	}
	return string(out)
}

// RandomLetter returns a random letter (a-z, A-Z)
func (faker *Faker) RandomLetter() rune {
	return letters[faker.random.Intn(len(letters))]
}

// Letterify replaces all question marks within a string with random letters.
func (faker *Faker) Letterify(s string) string {
	var out []rune
	for _, c := range s {
		if c == '?' {
			out = append(out, faker.RandomLetter())
		} else {
			out = append(out, c)
		}
	}
	return string(out)
}

// Bothify replaces all hash signs within a string with random digits and all
func (faker *Faker) Bothify(s string) string {
	var out []rune
	for _, c := range s {
		if c == '?' {
			out = append(out, faker.RandomLetter())
		} else if c == '#' {
			out = append(out, faker.RandomDigit())
		} else {
			out = append(out, c)
		}
	}
	return string(out)
}

func (faker *Faker) shuffleStrings(slice []string) {
	for i := range slice {
		j := faker.random.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Sample returns a sample (some random items) of a given array. The sample entries are not guaranteed to be unique.
//
// faker.Sample([]string{"a", "b", "c", "d", "e", "f"}, 6) might return []string{"e", "a", "e", "b", "f", "c"}
func (faker *Faker) Sample(arr []string, num int) []string {
	var result []string

	for i := 0; i < num; i++ {
		idx := faker.random.Intn(len(arr))
		result = append(result, arr[idx])
	}
	return result
}

// AllLocales returns a list of all available locales supported by gofaker.
func AllLocales() []string {
	return allLocales
}
