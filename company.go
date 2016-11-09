package gofaker

import (
	"fmt"
	"strconv"
	"strings"
)

type Company struct {
	faker *Faker
}

func (company *Company) Name() string {
	return company.faker.MustParse("company.name")
}

func (company *Company) Suffix() string {
	return company.faker.MustParse("company.suffix")
}

func (company *Company) Industry() string {
	return company.faker.MustParse("company.industry")
}

func (company *Company) Bullshit() string {
	lists := company.faker.MustFetchList("company.bs")
	var parts []string
	for _, p := range lists {
		parts = append(parts, company.faker.randomValue(p))
	}
	return strings.Join(parts, " ")
}

func (company *Company) EIN() string {
	return company.faker.Numerify("##-#######")
}

func (company *Company) DUNSNumber() string {
	return company.faker.Numerify("##-###-####")
}

func (company *Company) SwedishOrganisationNumber() string {
	base := company.faker.Numerify("#########")
	return fmt.Sprintf("%s%d", base, luhn(base))
}

func luhn(luhnString string) int {
	source := strings.Split(luhnString, "")
	checksum := 0
	double := false

	for i := len(source) - 1; i > -1; i-- {
		t, _ := strconv.ParseInt(source[i], 10, 8)
		n := int(t)

		if double {
			n = n * 2
		}
		double = !double

		if n >= 10 {
			n = n - 9
		}

		checksum += n
	}

	return checksum
}
