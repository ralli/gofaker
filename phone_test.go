package gofaker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhone(t *testing.T) {
	f, err := NewFaker("en-US")
	assert.NoError(t, err)

	t.Run("PhoneNumber", func(t *testing.T) { assert.NotEmpty(t, f.Phone.PhoneNumber()) })
	t.Run("CellPhone", func(t *testing.T) { assert.NotEmpty(t, f.Phone.CellPhone()) })
	t.Run("AreaCode", func(t *testing.T) { assert.NotEmpty(t, f.Phone.AreaCode()) })
	t.Run("ExchangeCode", func(t *testing.T) { assert.NotEmpty(t, f.Phone.ExchangeCode()) })
	t.Run("SubscriberNumber", func(t *testing.T) { assert.NotEmpty(t, f.Phone.SubscriberNumber()) })
}
