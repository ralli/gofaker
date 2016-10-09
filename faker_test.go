package gofaker

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFakerParse(t *testing.T) {
	f, err := NewFaker("en")
	assert.NoError(t, err)
	v, err := f.Parse("address.city_prefix")
	assert.NoError(t, err)
	fmt.Println(v)
	_, err = f.Parse("dummy")
	assert.Error(t, err)
}

func TestFakerRandomValue(t *testing.T) {
	f, err := NewFaker("en")
	assert.NoError(t, err)
	assert.Equal(t, "N/A", f.randomValue(nil))
}

func TestExtractSubKeys(t *testing.T) {
	const in = `#{city_prefix} #{Name.first_name}#{city_suffix}`
	a := extractSubKeys(in)
	assert.Equal(t, "[city_prefix Name.first_name city_suffix]", fmt.Sprint(a))
}
