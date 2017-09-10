package gofaker

import "math"

// Commerce provides functions to generate fake data for products sold in a web shop.
type Commerce struct {
	faker *Faker
}

// Color generates the name of a products color like "yellow".
func (commerce *Commerce) Color() string {
	return commerce.faker.MustParse("color.name")
}

// Department generates a category or department name like "Books", "Movies" or "Electronics".
func (commerce *Commerce) Department() string {
	return commerce.faker.MustParse("commerce.department")
}

// ProductName generates a random product name.
func (commerce *Commerce) ProductName() string {
	return commerce.faker.MustParse("commerce.product_name.adjective") + " " + commerce.faker.MustParse("commerce.product_name.material") + " " + commerce.faker.MustParse("commerce.product_name.product")
}

// Price generates a price in the range from 0.00 to 99.99.
func (commerce *Commerce) Price() float64 {
	return math.Floor(commerce.faker.random.Float64()*100*100) / 100.0
}
