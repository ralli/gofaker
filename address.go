package gofaker

import "fmt"

// Address does something
type Address struct {
	faker *Faker
}

func (a *Address) City() string {
	return a.faker.MustParse("address.city")
}

func (a *Address) StreetName() string {
	return a.faker.MustParse("address.street_name")
}

func (a *Address) SteetAddress() string {
	return a.faker.Numerify(a.faker.MustParse("address.street_address"))
}

func (a *Address) BuildingNumber() string {
	return a.faker.Bothify(a.faker.MustParse("address.building_number"))
}

func (a *Address) ZipCode() string {
	return a.faker.Bothify(a.faker.MustParse("address.postcode"))
}

func (a *Address) ZipCodeByState(stateName string) string {
	return a.faker.Bothify(a.faker.MustParse("address.postcode_by_state") + stateName)
}

func (a *Address) TimeZone() string {
	return a.faker.MustParse("address.time_zone")
}

func (a *Address) StreetSuffix() string {
	return a.faker.MustParse("address.street_suffix")
}

func (a *Address) CitySuffix() string {
	return a.faker.MustParse("address.city_suffix")
}

func (a *Address) CityPrefix() string {
	return a.faker.MustParse("address.city_prefix")
}

func (a *Address) StateAbbr() string {
	return a.faker.MustParse("address.state_abbr")
}

func (a *Address) State() string {
	return a.faker.MustParse("address.state")
}

func (a *Address) Country() string {
	return a.faker.MustParse("address.country")
}

func (a *Address) CountryCode() string {
	return a.faker.MustParse("address.country_code")
}

func (a *Address) Latitude() string {
	return fmt.Sprintf("%.2f", (a.faker.random.Float64()*180)-90)
}

func (a *Address) Longitude() string {
	return fmt.Sprintf("%.2f", (a.faker.random.Float64()*360)-180)
}
