package gofaker

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFaker(t *testing.T) {
	f, err := NewFaker("en")
	assert.NoError(t, err)

	t.Run("ParseSuccess", func(t *testing.T) {
		_, err := f.Parse("address.city_prefix")
		assert.NoError(t, err)
	})

	t.Run("ParseError", func(t *testing.T) {
		_, err := f.Parse("dummy")
		assert.Error(t, err)
	})

	t.Run("EmptyRandomValue", func(t *testing.T) {
		assert.Equal(t, "N/A", f.randomValue(nil))
	})

	t.Run("ParseWithSubKeys", func(t *testing.T) {
		a, err := f.Parse("address.city")
		assert.NoError(t, err)
		assert.NotContains(t, a, "#")
	})

	t.Run("Numerify", func(t *testing.T) {
		out := f.Numerify("###aaa#b&")
		re := regexp.MustCompile(`^\d\d\daaa\db&$`)
		assert.Regexp(t, re, out)
	})

	t.Run("Letterify", func(t *testing.T) {
		out := f.Letterify("???012?b&")
		re := regexp.MustCompile(`^[a-zA-Z]{3}012[a-zA-Z]b&$`)
		assert.Regexp(t, re, out)
	})

	t.Run("Bothify", func(t *testing.T) {
		out := f.Bothify("???###?b&")
		re := regexp.MustCompile(`^[a-zA-Z]{3}\d{3}[a-zA-Z]b&$`)
		assert.Regexp(t, re, out)
	})
}

func TestExtractSubKeys(t *testing.T) {
	const in = `#{city_prefix} #{Name.first_name}#{city_suffix}`
	a := extractSubKeys(in)
	assert.Equal(t, "[city_prefix Name.first_name city_suffix]", fmt.Sprint(a))
}

func TestNewFakerError(t *testing.T) {
	_, err := NewFaker("dummy")
	assert.Error(t, err)
}

type SnakeTest struct {
	input  string
	output string
}

var tests = []SnakeTest{
	{"a", "a"},
	{"snake", "snake"},
	{"A", "a"},
	{"ID", "id"},
	{"MOTD", "motd"},
	{"Snake", "snake"},
	{"SnakeTest", "snake_test"},
	{"SnakeID", "snake_id"},
	{"SnakeIDGoogle", "snake_id_google"},
	{"LinuxMOTD", "linux_motd"},
	{"OMGWTFBBQ", "omgwtfbbq"},
	{"omg_wtf_bbq", "omg_wtf_bbq"},
}

func TestToSnake(t *testing.T) {
	for _, test := range tests {
		if ToSnake(test.input) != test.output {
			t.Errorf(`ToSnake("%s"), wanted "%s", got \%s"`, test.input, test.output, ToSnake(test.input))
		}
	}
}

func TestBaseKeyPrefix(t *testing.T) {
	baseKey := "a.b.c.d"
	assert.Equal(t, "a.b.c", getBaseKeyPrefix(baseKey))
}