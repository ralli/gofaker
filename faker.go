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
	random  *rand.Rand
	data    Data
	Name    *Name
	Address *Address
}

var digits = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
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

// ToSnake convert the given string to snake case following the Golang format:
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

// NewFaker creates a new Faker instance for a given locale
func NewFaker(locale string) (*Faker, error) {
	data, err := NewData(locale)
	if err != nil {
		return nil, err
	}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	faker := &Faker{random: random, data: data}
	faker.Name = &Name{faker}
	faker.Address = &Address{faker}
	return faker, nil
}

func (faker *Faker) RandomDigit() rune {
	return digits[faker.random.Intn(10)]
}

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

func (faker *Faker) RandomLetter() rune {
	return letters[faker.random.Intn(len(letters))]
}

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
