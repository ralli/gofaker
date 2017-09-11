package gofaker

// App provides functions related to apps.
type App struct {
	faker *Faker
}

// Name generates the apps name.
func (a *App) Name() string {
	return a.faker.MustParse("app.name")
}

// Version generates the apps version.
func (a *App) Version() string {
	return a.faker.Bothify(a.faker.MustParse("app.version"))
}

// Author generates the apps authors name.
func (a *App) Author() string {
	return a.faker.MustParse("app.author")
}
