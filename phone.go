package gofaker

// Phone provides functions to generate phone numbers.
type Phone struct {
	faker *Faker
}

// PhoneNumber generates a phone number. Different phone numbers will be generated in different formats. The formats are locale dependent.
func (phone *Phone) PhoneNumber() string {
	return phone.faker.Numerify(phone.faker.MustFetch("phone_number.formats"))
}

// CellPhone genrates valid cell phone numbers.
func (phone *Phone) CellPhone() string {
	return phone.faker.Numerify(phone.faker.MustFetch("cell_phone.formats"))
}

// SubscriberNumber generates subscriber numbers.
func (phone *Phone) SubscriberNumber() string {
	n := phone.faker.random.Intn(2) + 3
	var out []rune
	for i := 0; i < n; i++ {
		out = append(out, phone.faker.RandomDigit())
	}
	return string(out)
}
