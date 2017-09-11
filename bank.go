package gofaker

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode"
)

// Bank provides method to generate bank related test data.
type Bank struct {
	faker *Faker
}

// Name generates a banks name.
func (b *Bank) Name() string {
	return b.faker.MustFetch("bank.name")
}

// BIC generates a banks BIC (bank identifier code).
func (b *Bank) BIC() string {
	return b.faker.MustFetch("bank.swift_bic")
}

func (b *Bank) IBAN() string {
	iban, err := b.NationalIBAN("DE")
	if err != nil {
		panic(err)
	}
	return iban
}

func getIBANInfo(ibanDetails []map[string]string, countryCode string) map[string]string {
	for _, m := range ibanDetails {
		if m["bank_country_code"] == countryCode {
			return m
		}
	}
	return nil
}

// NationalIBAN generates a IBAN for the country given. Returns an error if there are no valid IBANS for the country.
func (b *Bank) NationalIBAN(countryCode string) (string, error) {
	details := b.faker.data.GetStringMapList("bank.iban_details")
	m := getIBANInfo(details, countryCode)
	if m == nil {
		return "", fmt.Errorf("invalid country code for iban '%s'", countryCode)
	}
	s := bytes.NewBufferString(countryCode)
	numLetters, err := strconv.Atoi(m["iban_letter_code"])
	if err != nil {
		return "", err
	}
	for i := 0; i < numLetters; i++ {
		s.WriteRune(unicode.ToUpper(b.faker.RandomLetter()))
	}
	numDigits, err := strconv.Atoi(m["iban_digits"])
	if err != nil {
		return "", err
	}
	for i := 0; i < numDigits; i++ {
		s.WriteRune(b.faker.RandomDigit())
	}
	return s.String(), nil
}
