package gofaker

import "fmt"

// Code provides functions to generate codes (ISBN numbers and the like).
type Code struct {
	faker *Faker
}

// Base10ISBN generates an old (10 digit) ISBN number (international standard book number).
// Users normally would like to generate a Base13ISBN. Since 2007 all ISBNs have to be base 13.
func (code *Code) Base10ISBN() string {
	num := code.faker.Numerify("#########")
	sum := 0
	for i, v := range num {
		val := int(v - '0')
		idx := i + 1
		sum += idx * val
	}
	sum %= 11
	var remainder string
	if sum == 10 {
		remainder = "X"
	} else {
		remainder = fmt.Sprint(sum)
	}

	return num + remainder
}

// Base10ISBN generates an ISBN number (international standard book number).
func (code *Code) Base13ISBN() string {
	num := code.faker.Numerify("############")
	sum := 0
	for i, v := range num {
		val := int(v - '0')
		idx := i + 1
		if idx%2 == 0 {
			sum += val
		} else {
			sum += val * 3
		}
	}

	sum %= 10
	check := (10 - sum) % 10
	remainder := fmt.Sprint(check)
	return num + remainder
}
