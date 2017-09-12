package gofaker

// Food provides functions to generate fake data related to food.
type Food struct {
	faker *Faker
}

// Dish generates a dishes name.
func (f *Food) Dish() string {
	return f.faker.MustFetch("food.dish")
}

// Ingredient generates an indigrients name.
func (f *Food) Ingredient() string {
	return f.faker.MustFetch("food.ingredients")
}

// Spice generates a spices name.
func (f *Food) Spice() string {
	return f.faker.MustFetch("food.spices")
}

// Measurement generates a mesurement like "1 teaspoon".
func (f *Food) Measurement() string {
	return f.faker.MustFetch("food.measurement_sizes") + " " + f.faker.MustFetch("food.measurements")
}

// MetricMeasurement generates a metric mesurement like "litre".
func (f *Food) MetricMeasurement() string {
	return f.faker.MustFetch("food.metric_measurements")
}
