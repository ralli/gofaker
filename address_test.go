package gofaker

import (
	"strconv"
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestAddress(t *testing.T) {
	f, err := NewFaker("en")
	assert.NoError(t, err)
	t.Run("City", func(t *testing.T) { assert.NotEmpty(t, f.Address.City()) })
	t.Run("StreetName", func(t *testing.T) { assert.NotEmpty(t, f.Address.StreetName()) })
	t.Run("StreetAddress", func(t *testing.T) { assert.NotEmpty(t, f.Address.StreetAddress()) })
	t.Run("BuildingNumber", func(t *testing.T) { assert.NotEmpty(t, f.Address.BuildingNumber()) })
	t.Run("ZipCode", func(t *testing.T) { assert.NotEmpty(t, f.Address.ZipCode()) })
	t.Run("ZipCodeByState", func(t *testing.T) { assert.NotEmpty(t, f.Address.ZipCodeByState("AZ")) })
	t.Run("TimeZone", func(t *testing.T) { assert.NotEmpty(t, f.Address.TimeZone()) })
	t.Run("StateAbbr", func(t *testing.T) { assert.NotEmpty(t, f.Address.StateAbbr()) })
	t.Run("State", func(t *testing.T) { assert.NotEmpty(t, f.Address.State()) })
	t.Run("Country", func(t *testing.T) { assert.NotEmpty(t, f.Address.Country()) })
	t.Run("CountryCode", func(t *testing.T) { assert.NotEmpty(t, f.Address.CountryCode()) })
	t.Run("Latitude", func(t *testing.T) {
		out := f.Address.Latitude()
		lat, err := strconv.ParseFloat(out, 64)
		assert.NoError(t, err)
		if lat > 90.0 {
			t.Errorf("expected %f to be less than 90.0", lat)
		}
		if lat < -90.0 {
			t.Errorf("expected %f to be greater or equal than -90.0", lat)
		}
	})

	t.Run("Longitude", func(t *testing.T) {
		out := f.Address.Longitude()
		lon, err := strconv.ParseFloat(out, 64)
		assert.NoError(t, err)
		if lon > 180.0 {
			t.Errorf("expected %f to be less than 180.0", lon)
		}
		if lon < -180.0 {
			t.Errorf("expected %f to be greater or equal than -180.0", lon)
		}
	})
}

func ExampleAddress_City() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Address.City())
	// Output: West Camilla
}

func ExampleAddress_StreetName() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Address.StreetName())
	// Output: Willms Grove
}

func ExampleAddress_StreetAddress() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Address.StreetAddress())
	// Output: 576 Maybelle Prairie
}

func ExampleAddress_BuildingNumber() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Address.BuildingNumber())
	// The building number may actually have a number of different formats
	// Output: 780
}

func ExampleAddress_ZipCode() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Address.ZipCode())
	// Output: 78035-7683
}

func ExampleAddress_ZipCodeByState() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Address.ZipCodeByState("AZ"))
	// Output: 85078
}

func ExampleAddress_TimeZone() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Address.TimeZone())
	// Output: Europe/Dublin
}

func ExampleAddress_StateAbbr() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Address.StateAbbr())
	// Output: CO
}

func ExampleAddress_State() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Address.State())
	// Output: Colorado
}

func ExampleAddress_Country() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Address.Country())
	// Output: Tanzania
}

func ExampleAddress_CountryCode() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Address.CountryCode())
	// Output: CY
}

func ExampleAddress_Community() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Address.Community())
	// Output: Willow Court
}

func ExampleAddress_FullAddress() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Address.FullAddress())
	// Output: Suite 823 21542 Schaefer Path, New Lukasbury, LA 44541-2979
}

func ExampleAddress_Longitude() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Address.Longitude())
	// Output: -45.71
}

func ExampleAddress_Latitude() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Address.Latitude())
	// Output: -22.85
}
