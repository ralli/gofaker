package gofaker

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestPhone(t *testing.T) {
	f, err := NewFaker("en-US")
	assert.NoError(t, err)

	t.Run("PhoneNumber", func(t *testing.T) { assert.NotEmpty(t, f.Phone.PhoneNumber()) })
	t.Run("CellPhone", func(t *testing.T) { assert.NotEmpty(t, f.Phone.CellPhone()) })
	t.Run("SubscriberNumber", func(t *testing.T) { assert.NotEmpty(t, f.Phone.SubscriberNumber()) })
}

func ExamplePhone_PhoneNumber() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Phone.PhoneNumber())
	// Output: (780) 357-6839
}

func ExamplePhone_Cellphone() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Phone.CellPhone())
	// Output: (780) 357-6839
}

func ExamplePhone_SubscriberNumber() {
	f, _ := NewFaker("en")
	f.Reset()
	fmt.Println(f.Phone.SubscriberNumber())
	// Output: 7803
}
