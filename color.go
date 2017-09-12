package gofaker

// Color provides functions related to colors.
type Color struct {
	faker *Faker
}

// Name returns a colors name.
func (c *Color) Name() string {
	return c.faker.MustFetch("color.name")
}
