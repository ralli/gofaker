package gofaker

type Compass struct {
	faker *Faker
}

func (c *Compass) Cardinal() string {
	return c.faker.MustFetch("compass.cardinal.word")
}

func (c *Compass) Ordinal() string {
	return c.faker.MustFetch("compass.ordinal.word")
}

func (c *Compass) HalfWind() string {
	return c.faker.MustFetch("compass.half-wind.word")
}

func (c *Compass) QuarterWind() string {
	return c.faker.MustFetch("compass.quarter-wind.word")
}

func (c *Compass) Direction() string {
	return c.faker.MustParse("compass.direction")
}

func (c *Compass) Abbreviation() string {
	return c.faker.MustParse("compass.abbreviation")
}

func (c *Compass) Azimuth() string {
	return c.faker.MustParse("compass.azimuth")
}

func (c *Compass) CardinalAbbreviation() string {
	return c.faker.MustFetch("compass.cardinal.abbreviation")
}

func (c *Compass) OrdinalAbbreviation() string {
	return c.faker.MustFetch("compass.ordinal.abbreviation")
}

func (c *Compass) HalfWindAbbreviation() string {
	return c.faker.MustFetch("compass.half-wind.abbreviation")
}

func (c *Compass) QuarterWindAbbreviation() string {
	return c.faker.MustFetch("compass.quarter-wind.abbreviation")
}

func (c *Compass) CardinalAzimuth() string {
	return c.faker.MustFetch("compass.cardinal.azimuth")
}

func (c *Compass) OrdinalAzimuth() string {
	return c.faker.MustFetch("compass.ordinal.azimuth")
}

func (c *Compass) HalfWindAzimuth() string {
	return c.faker.MustFetch("compass.half-wind.azimuth")
}

func (c *Compass) QuarterWindAzimuth() string {
	return c.faker.MustFetch("compass.quarter-wind.azimuth")
}
