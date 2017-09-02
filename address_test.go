package gofaker

import (
	"regexp"
	"strconv"
	"testing"

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
	t.Run("ZipCodeByState", func(t *testing.T) {
		out := f.Address.ZipCodeByState("AZ")
		assert.Regexp(t, regexp.MustCompile(".*AZ$"), out)
	})
	t.Run("TimeZone", func(t *testing.T) { assert.NotEmpty(t, f.Address.TimeZone()) })

	t.Run("StreetSuffix", func(t *testing.T) { assert.NotEmpty(t, f.Address.StreetSuffix()) })
	t.Run("CitySuffix", func(t *testing.T) { assert.NotEmpty(t, f.Address.CitySuffix()) })
	t.Run("CityPrefix", func(t *testing.T) { assert.NotEmpty(t, f.Address.CityPrefix()) })
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
