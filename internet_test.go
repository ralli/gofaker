package gofaker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternet(t *testing.T) {
	const userNameRegex = `\w+[._-]\w+`
	const domainNameRegex = `^[a-z][a-z_-]*`

	f, err := NewFaker("en")
	assert.NoError(t, err)

	t.Run("UserWithName", func(t *testing.T) {
		assert.Regexp(t, "ralph[_.]juhnke", f.Internet.UserWithName("Ralph Juhnke"))
	})

	t.Run("UserName", func(t *testing.T) {
		userName := f.Internet.UserName()
		assert.Regexp(t, userNameRegex, userName)
	})

	t.Run("PasswordWithLength", func(t *testing.T) {
		assert.NotEmpty(t, f.Internet.PasswordWithLength(5, 6))
	})

	t.Run("Password", func(t *testing.T) {
		assert.NotEmpty(t, f.Internet.Password())
	})

	t.Run("DomainWord", func(t *testing.T) {
		assert.Regexp(t, domainNameRegex, f.Internet.DomainWord())
	})

	t.Run("EmailWithUsername", func(t *testing.T) {
		assert.Regexp(t, "^ralph.juhnke@[a-z][a-z_-]*\\.\\w+$", f.Internet.EmailWithUsername("Ralph Juhnke"))
	})

	t.Run("Email", func(t *testing.T) {
		assert.Regexp(t, "^\\w+[._-]\\w+@[a-z][a-z_-]*\\.\\w+$", f.Internet.Email())
	})

	t.Run("FreeMailWithName", func(t *testing.T) {
		assert.Regexp(t, "^ralph.juhnke@[a-z][a-z_-]*\\.\\w+$", f.Internet.FreeMailWithName("Ralph Juhnke"))
	})

	t.Run("FreeMail", func(t *testing.T) {
		assert.Regexp(t, "^\\w+[._-]\\w+@[a-z][a-z_-]*\\.\\w+$", f.Internet.FreeMail())
	})

	t.Run("SafeMailWithName", func(t *testing.T) {
		assert.Regexp(t, "^ralph.juhnke@example\\.(org|com|net)$", f.Internet.SafeMailWithName("Ralph Juhnke"))
	})

	t.Run("SafeMailWithName", func(t *testing.T) {
		assert.Regexp(t, "^\\w+[._-]\\w+@example\\.(org|com|net)$", f.Internet.SafeMail())
	})

	t.Run("MacAddress", func(t *testing.T) {
		assert.Regexp(t, `^[a-f0-9]{2}:[a-f0-9]{2}:[a-f0-9]{2}:[a-f0-9]{2}:[a-f0-9]{2}:[a-f0-9]{2}$`, f.Internet.MacAddress())
	})

	t.Run("IPv4Address", func(t *testing.T) {
		assert.Regexp(t, `^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`, f.Internet.IPv4Address())
	})
}
