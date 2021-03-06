package gofaker

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleCode_ASIN() {
	faker, _ := NewFaker("en")
	faker.Reset()
	fmt.Println(faker.Code.ASIN())
	// Output: B000654P8C
}

func TestCodeBase10ISBN(t *testing.T) {
	faker, err := NewFaker("en")
	assert.NoError(t, err)
	faker.Reset()
	out := faker.Code.Base10ISBN()
	assert.Equal(t, "5780357684", out)
	for i := 0; i < 100; i++ {
		out = faker.Code.Base10ISBN()
		assert.Equal(t, 10, len(out))
	}
}

func TestBase13ISBN(t *testing.T) {
	faker, err := NewFaker("en")
	assert.NoError(t, err)
	faker.Reset()
	out := faker.Code.Base13ISBN()
	assert.Equal(t, "5780357683972", out)
}
