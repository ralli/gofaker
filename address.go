package gofaker

import "fmt"

// Address provides functions to generate postal address related fake data.
type Address struct {
	faker *Faker
}

// City generates a cities name.
func (a *Address) City() string {
	return a.faker.MustParse("address.city")
}

// StreetName generates a street name.
func (a *Address) StreetName() string {
	return a.faker.MustParse("address.street_name")
}

// StreetAddress generates a street name with locale dependent additional data (a building number for example).
func (a *Address) StreetAddress() string {
	return a.faker.Numerify(a.faker.MustParse("address.street_address"))
}

// BuildingNumber generates a building number.
func (a *Address) BuildingNumber() string {
	return a.faker.Bothify(a.faker.MustParse("address.building_number"))
}

// ZipCode generates a zip code.
func (a *Address) ZipCode() string {
	return a.faker.Bothify(a.faker.MustParse("address.postcode"))
}

// ZipCodeByState generates a zipcode with an attached state name.
func (a *Address) ZipCodeByState(stateName string) string {
	return a.faker.Bothify(a.faker.MustParse("address.postcode_by_state." + stateName))
}

// TimeZone generates a time zone name like "Europe/Berlin".
func (a *Address) TimeZone() string {
	return a.faker.MustParse("address.time_zone")
}

// StateAbbr generates a abbreviated state name like "AZ" for "Arizona".
func (a *Address) StateAbbr() string {
	return a.faker.MustParse("address.state_abbr")
}

// State generates a states name like "Arizona".
func (a *Address) State() string {
	return a.faker.MustParse("address.state")
}

// Country generates a countries name like "Bulgaria".
func (a *Address) Country() string {
	return a.faker.MustParse("address.country")
}

// CountryCode generates a two letter ISO Country Code like "DE" for "Germany".
func (a *Address) CountryCode() string {
	return a.faker.MustParse("address.country_code")
}

func (a *Address) Community() string {
	return a.faker.MustParse("address.community")
}

func (a *Address) FullAddress() string {
	return a.faker.Bothify(a.faker.MustParse("address.full_address"))
}

// Latitude generates the latitude of a geographical position.
func (a *Address) Latitude() string {
	return fmt.Sprintf("%.2f", (a.faker.random.Float64()*180)-90)
}

// Longitude generates the longitude of a geographical position.
func (a *Address) Longitude() string {
	return fmt.Sprintf("%.2f", (a.faker.random.Float64()*360)-180)
}
