package gofaker

import "math"

type Commerce struct {
	faker *Faker
}

func (commerce *Commerce) Color() string {
	return commerce.faker.MustParse("color.name")
}

func (commerce *Commerce) Department() string {
	return commerce.faker.MustParse("commerce.department")
}

func (commerce *Commerce) ProductName() string {
	return commerce.faker.MustParse("commerce.product_name.adjective") + " " + commerce.faker.MustParse("commerce.product_name.material") + " " + commerce.faker.MustParse("commerce.product_name.product")
}

func (commerce *Commerce) Price() float64 {
	return math.Floor(commerce.faker.random.Float64()*100*100) / 100.0
}
