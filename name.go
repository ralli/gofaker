package gofaker

// Name generates name related fake data
type Name struct {
	faker *Faker
}

// FirstName generates a first name
func (name *Name) FirstName() string {
	return name.faker.MustParse("name.first_name")
}

// LastName generates a last name
func (name *Name) LastName() string {
	return name.faker.MustParse("name.last_name")
}

// Prefix generates a name prefix
func (name *Name) Prefix() string {
	return name.faker.MustParse("name.prefix")
}

func (name *Name) Suffix() string {
	return name.faker.MustParse("name.suffix")
}

// Title returns something like "Senior Vice Chief"
func (name *Name) Title() string {
	descriptor := name.faker.MustParse("name.title.descriptor")
	level := name.faker.MustParse("name.title.level")
	job := name.faker.MustParse("name.title.job")
	return descriptor + " " + level + " " + job
}
