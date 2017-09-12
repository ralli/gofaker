package gofaker

type Dessert struct {
	faker *Faker
}

func (d *Dessert) Variety() string {
	return d.faker.MustFetch("dessert.variety")
}

func (d *Dessert) Topping() string {
	return d.faker.MustFetch("dessert.topping")
}

func (d *Dessert) Flavor() string {
	return d.faker.MustFetch("dessert.flavor")
}
