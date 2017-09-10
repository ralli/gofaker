package gofaker

import (
	"fmt"
	"log"
)

type Person struct {
	FirstName string
	LastName  string
	Login     string
	Email     string
}

func Example() {
	// an error may occur if a non existing locale is provided
	faker, err := NewFaker("en")
	if err != nil {
		log.Fatal(err)
	}

	// Reset sets the internally used random number generator to a fixed seed.
	// This ensures the same fake data is generated on subsequent calls.
	// This becomes handy in unit tests and the like.

	faker.Reset()

	person := Person{FirstName: faker.Name.FirstName(), LastName: faker.Name.LastName()}
	person.Login = faker.Internet.UserWithName(person.FirstName + " " + person.LastName)
	person.Email = faker.Internet.SafeMailWithName(person.FirstName + " " + person.LastName)
	fmt.Printf("%s: %s %s (%s)", person.Login, person.FirstName, person.LastName, person.Email)

	// Output: tommie_willms: Tommie Willms (tommie_willms@example.com)
}
