package gofaker

import "fmt"

func ExampleBank_Name() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Bank.Name())
	// Output: ABN AMRO MEZZANINE (UK) LIMITED
}

func ExampleBank_BIC() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Bank.BIC())
	// Output: ABAZGB21
}

func ExampleBank_IBAN() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Bank.IBAN())
	// Output: DE57803576839758232154
}

func ExampleBank_NationalIBAN() {
	f, _ := NewFaker("en")
	f.Reset()
	iban, _ := f.Bank.NationalIBAN("BG")
	fmt.Println(iban)
	// Output: BGHRUK35768397582321
}
