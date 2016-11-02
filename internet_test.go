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
	t.Run("PasswordWithLength", func(t *testing.T) {
		assert.NotEmpty(t, f.Internet.PasswordWithLength(5, 6))
	})
	t.Run("Password", func(t *testing.T) {
		assert.NotEmpty(t, f.Internet.Password())
	})
	t.Run("DomainWord", func(t *testing.T) {
		assert.NotEmpty(t, f.Internet.DomainWord())
	})
	t.Run("EmailWithUsername", func(t *testing.T) {
		assert.NotEmpty(t, f.Internet.EmailWithUsername("Ralph Juhnke"))
	})
	t.Run("Email", func(t *testing.T) {
		assert.NotEmpty(t, f.Internet.Email())
	})
	t.Run("FreeMailWithName", func(t *testing.T) {
		assert.NotEmpty(t, f.Internet.FreeMailWithName("Ralph Juhnke"))
	})

}
