package gofaker

import "fmt"

func ExampleCompanyName() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Company.Name())
	// Output: Willms, Crona and Swift
}

func ExampleCompanySuffix() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Company.Suffix())
	// Output: and Sons
}

func ExampleCompanyIndustry() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Company.Industry())
	// Output: Research
}

func ExampleCompanyEIN() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Company.EIN())
	// Output: 57-8035768
}

func ExampleCompanyDUNSNumber() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Company.DUNSNumber())
	// Output: 57-803-5768
}

func ExampleLuhn() {
	in := "446667651"
	out := luhn(in)
	fmt.Println(out)
	// Output: 40
}

func ExampleSwedishOrganisationNumber() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Company.SwedishOrganisationNumber())
	// Output: 57803576840
}
