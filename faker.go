package gofaker

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

// Faker does something
type Faker struct {
	random *rand.Rand
	data   Data
}

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

// Parse gets and expands a key
func (faker *Faker) Parse(key string) (string, error) {
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
	return &Faker{random, data}, nil
}
