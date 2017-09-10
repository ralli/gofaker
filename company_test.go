package gofaker

import "fmt"

func ExampleCompany_Name() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Company.Name())
	// Output: Willms, Crona and Swift
}

func ExampleCompany_Bullshit() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Company.Bullshit())
	// Output: evolve global solutions
}

func ExampleCompany_Industry() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Company.Industry())
	// Output: Research
}

func ExampleCompany_EIN() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Company.EIN())
	// Output: 57-8035768
}

func ExampleCompany_DUNSNumber() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Company.DUNSNumber())
	// Output: 57-803-5768
}

func ExampleCompany_SwedishOrganisationNumber() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Company.SwedishOrganisationNumber())
	// Output: 57803576840
}
