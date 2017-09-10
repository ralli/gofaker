// Package gofaker helps generating test data like persons names, addresses, email addresses, etc.
//
// Credits
//
// gofaker is (more or less) a port of Benjamin Curtis (stympy) faker library to generate
// fake data (which is itself a port of Perl's Data::Faker library).
//
// Locales
//
// Faker supports generating fake data in different languages. You can find a list of all available locales in the example of the function AllLocales().
//
// Thread safety
//
// Currently Faker objects are not thread safe since they use a single random number generator instance. Genrating
// fake data from multiple goroutines concurrently is still possible when using a separate Faker instance per goroutine.
package gofaker
