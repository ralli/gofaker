package data

import (
	"testing"

	yaml "gopkg.in/yaml.v2"

	"github.com/stretchr/testify/assert"
)

func TestAssetDir(t *testing.T) {
	data, err := AssetDir("data")
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
}

func TestEnYml(t *testing.T) {
	_, err := Asset("data/en.yml")
	assert.NoError(t, err)
}

func TestParseData(t *testing.T) {
	data, err := Asset("data/en.yml")
	assert.NoError(t, err)
	var yamlData interface{} = make(map[string]interface{})
	err = yaml.Unmarshal(data, &yamlData)
	assert.NoError(t, err)
	m := yamlData.(map[interface{}]interface{})
	assert.NotNil(t, m["en"])
}
