package gofaker

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"
)

var testData = []byte(`en:
  faker:
    separator: ' & '
    address:
      city_prefix: [North, East, West, South, New, Lake, Port]`)

var existingInputs = []struct {
	in  string
	out string
}{
	{"en.faker.address.city_prefix", "[North East West South New Lake Port]"},
	{"en.faker.separator", "[ & ]"},
}

func TestKeys(t *testing.T) {
	in := "a.b.cc"
	out := keys(in)
	assert.Equal(t, "[a b cc]", fmt.Sprint(out))
}

func TestLocaleDataGet(t *testing.T) {
	var yamlData interface{} = make(map[string]interface{})
	err := yaml.Unmarshal(testData, &yamlData)
	assert.NoError(t, err)
	localeData := &localeData{"en", yamlData}
	for _, m := range existingInputs {
		a := localeData.Get(m.in)
		assert.NotEmpty(t, a)
		assert.Equal(t, m.out, fmt.Sprint(a))
	}
}

func TestLocaleGetNil(t *testing.T) {
	var yamlData interface{} = make(map[string]interface{})
	err := yaml.Unmarshal(testData, &yamlData)
	assert.NoError(t, err)
	localeData := &localeData{"en", yamlData}
	nonExistingKeys := []string{"en", "en.faker", "en.faker.address", "en.faker.address.dummy"}
	for _, k := range nonExistingKeys {
		a := localeData.Get(k)
		assert.Empty(t, a)
	}
}

func TestLocaleNames(t *testing.T) {
	testInput := []struct {
		in  string
		out string
	}{
		{"pt-BR", "[pt-BR pt en]"},
		{"de", "[de en]"},
		{"en", "[en]"},
	}
	for _, v := range testInput {
		a := getFallbackLocales(v.in)
		s := fmt.Sprint(a)
		assert.Equal(t, v.out, s)
	}
}

func TestFakerDataGet(t *testing.T) {
	d, err := newData("de")
	assert.NoError(t, err)
	a := d.Get("address.city_prefix")
	assert.NotEmpty(t, a)
	assert.Equal(t, "[Nord Ost West Süd Neu Alt Bad Groß Klein Schön]", fmt.Sprint(a))
	a = d.Get("separator")
	assert.NotEmpty(t, a)
	assert.Equal(t, "[ & ]", fmt.Sprint(a))
}

func TestNewDataError(t *testing.T) {
	_, err := newData("giugu")
	assert.Error(t, err)
}
