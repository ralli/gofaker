package gofaker

import (
	"fmt"
	"testing"
	"testing/quick"
)

func TestCreditCard_ExpiryDate(t *testing.T) {
	faker, _ := NewFaker("en")
	f := func(x int) bool {
		date := faker.CreditCard.ExpiryDate()
		return date.Day() == 1
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func ExampleCreditCard_ExpiryDate() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.CreditCard.ExpiryDate().Format("2006-01-02"))
	// Outputs something like: 2020-04-01
}
