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

	t.Run("PrivateIPv4Address", func(t *testing.T) {
		assert.Regexp(t, `^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`, f.Internet.PrivateIPv4Address())
	})

	t.Run("isPrivateNet", func(t *testing.T) {
		in := []struct {
			input    string
			expected bool
		}{
			{"192.168.1.7", true},
			{"10.0.0.1", true},
			{"169.254.1.7", true},
			{"172.28.1.7", true},
			{"172.172.1.3", false},
			{"18.31.1.3", false},
		}

		t.Run("IPv6Address", func(t *testing.T) {
			assert.Regexp(t, `^[0-9a-f]{1,4}:[0-9a-f]{1,4}:[0-9a-f]{1,4}:[0-9a-f]{1,4}:[0-9a-f]{1,4}:[0-9a-f]{1,4}:[0-9a-f]{1,4}:[0-9a-f]{1,4}$`, f.Internet.IPv6Address())
		})

		for _, v := range in {
			if isPrivateNet(v.input) != v.expected {
				t.Errorf("expected isPrivateNet(%s) to be %v but was %v", v.input, v.expected, !v.expected)
			}
		}
	})
}
