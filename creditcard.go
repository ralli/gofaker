package gofaker

import (
	"bytes"
	"fmt"
	"strings"
	"time"
	"unicode"
)

type CreditCard struct {
	faker *Faker
}

var creditCardTypes = []string{"visa", "mastercard", "discover", "american_express",
	"diners_club", "jcb", "switch", "solo", "dankort",
	"maestro", "forbrugsforeningen", "laser"}

func onlyDigits(s string) string {
	var buf bytes.Buffer
	for _, r := range s {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

// Types returns a list of all credit card types.
func (f *CreditCard) Types() []string {
	return creditCardTypes
}

// TypeOf Returns one of the types given.
func (f *CreditCard) TypeOf(creditCardTypes []string) string {
	return f.faker.randomValue(creditCardTypes)
}

// Number returns either a visa or mastercard credit card number.
func (f *CreditCard) Number() string {
	return f.NumberOfType(f.TypeOf([]string{"visa", "mastercard"}))
}

// Number of type returns a credit card number of a given credit card type.A
func (f *CreditCard) NumberOfType(creditCardType string) string {
	key := "finance.credit_card." + creditCardType
	v := f.faker.Regexify(f.faker.Numerify(f.faker.MustFetch(key)))
	luhnDigit := fmt.Sprint(luhn(onlyDigits(v)))
	return strings.Replace(v, "L", luhnDigit, 1)
}

func firstDayOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// ExpiryDate returns a credit cards expiry date.
func (f *CreditCard) ExpiryDate() time.Time {
	hours := time.Duration(f.faker.random.Intn(4 * 365 * 24))
	return firstDayOfMonth(time.Now().Add(hours * time.Hour))
}

// ExpiryDateString returns the expiry date in the format MM/YY.
func (f *CreditCard) ExpiryDateString() string {
	return f.ExpiryDate().Format("01/06")
}
