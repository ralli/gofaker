package gofaker

// Name generates fake data related to persons names.
type Name struct {
	faker *Faker
}

// FirstName generates a first name.
func (name *Name) FirstName() string {
	return name.faker.MustParse("name.first_name")
}

// LastName generates a last name.
func (name *Name) LastName() string {
	return name.faker.MustParse("name.last_name")
}

// Name returns a full persons name.
func (name *Name) Name() string {
	return name.faker.MustParse("name.name")
}

// Title returns a job title like "Senior Vice Chief".
func (name *Name) Title() string {
	descriptor := name.faker.MustParse("name.title.descriptor")
	level := name.faker.MustParse("name.title.level")
	job := name.faker.MustParse("name.title.job")
	return descriptor + " " + level + " " + job
}
