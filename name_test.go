package gofaker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	f, err := NewFaker("de")
	assert.NoError(t, err)
	t.Run("FirstName", func(t *testing.T) {
		assert.NotEmpty(t, f.Name.FirstName())
	})
	t.Run("LastName", func(t *testing.T) {
		assert.NotEmpty(t, f.Name.LastName())
	})
	t.Run("Name", func(t *testing.T) {
		assert.NotEmpty(t, f.Name.Name())
	})
	t.Run("Prefix", func(t *testing.T) {
		assert.NotEmpty(t, f.Name.Prefix())
	})
	t.Run("Suffix", func(t *testing.T) {
		assert.NotEmpty(t, f.Name.Suffix())
	})
	t.Run("Title", func(t *testing.T) {
		assert.NotEmpty(t, f.Name.Title())
	})
	
}
