package gofaker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternet(t *testing.T) {
	f, err := NewFaker("en")
	assert.NoError(t, err)
	t.Run("UserWithName", func(t *testing.T) {
		assert.Regexp(t, "ralph[_.]juhnke", f.Internet.UserWithName("Ralph Juhnke"))
	})
	t.Run("UserName", func(t *testing.T) {
		assert.NotEmpty(t, f.Internet.UserName())
	})
}
