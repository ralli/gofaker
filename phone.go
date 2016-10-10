package gofaker

type Phone struct {
	faker *Faker
}

func (phone *Phone) PhoneNumber() string {
	return phone.faker.Numerify(phone.faker.MustFetch("phone_number.formats"))
}

func (phone *Phone) CellPhone() string {
	return phone.faker.Numerify(phone.faker.MustFetch("cell_phone.formats"))
}

func (phone *Phone) AreaCode() string {
	return phone.faker.FetchOrEmpty("phone_number.area_code")
}

func (phone *Phone) ExchangeCode() string {
	return phone.faker.FetchOrEmpty("phone_number.exchange_code")
}

func (phone *Phone) SubscriberNumber() string {
	n := phone.faker.random.Intn(2) + 3
	var out []rune
	for i := 0; i < n; i++ {
		out = append(out, phone.faker.RandomDigit())
	}
	return string(out)
}
