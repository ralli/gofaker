package gofaker

// Cat provides functions to genrate fake data related to cats.
type Cat struct {
	faker *Faker
}

// Name generates the cats name.
func (c *Cat) Name() string {
	return c.faker.MustFetch("cat.name")
}

// Breed generates the cats breed like "British Semipi-longhair".
func (c *Cat) Breed() string {
	return c.faker.MustFetch("cat.breed")
}

// Registry generates the cats registries name like "American Cat Fanciers Association".
func (c *Cat) Registry() string {
	return c.faker.MustFetch("cat.registry")
}
