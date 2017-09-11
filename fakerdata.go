package gofaker

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ralli/gofaker/data"
	yaml "gopkg.in/yaml.v2"
)

// Data provides access to fake data definitions used to make up fake data.
type fakerData interface {
	// Get returns a list of possible values for a given key. For example Get("name.first_name") will
	// return a list of valid first names.
	Get(key string) []string

	// GetList returns a list of lists of valid values for a given key. This is used for "bullshit" like sentences
	// which are built of a valid first, second and third word (subject, verb and object).
	// An example for a "bullshit" sentence is "implement value-added web-readyness".
	GetList(key string) [][]string

	// GetStringMapList returns a list of string maps. This method is used if structured data is needed to generate data.
	// An example is "bank.iban_details" where you need a country code, an iban letter code an a maximum length to generate
	// a valid IBAN.
	GetStringMapList(key string) []map[string]string
}

type localeData struct {
	locale string
	data   interface{}
}

// data holds Data objects for multiple locales.
// The Get operation will try all these locale data objects until a match is found.
type allData struct {
	locales []*localeData
}

var allLocales []string

var allDatas = make(map[string]*allData)

// keys returns an array of sub keys from a given key.
// ex. keys("address.prefix") returns ["address", "prefix"]
func keys(key string) []string {
	return strings.Split(key, ".")
}

// loadAllLocales loads a list of all available locales
// as a sorted list of strings
func loadAllLocales() []string {
	data, err := data.AssetDir("data")
	if err != nil {
		panic(err)
	}
	var result []string
	for _, v := range data {
		if strings.HasSuffix(v, ".yml") {
			result = append(result, strings.TrimSuffix(v, ".yml"))
		}
	}
	sort.Strings(result)
	return result
}

// existsLocale returns true, if a given locale exists
func existsLocale(locale string) bool {
	i := sort.SearchStrings(allLocales, locale)
	return i < len(allLocales)
}

// localeNames returns a list of fallback locales for a given locale.
// localeNames("pt_BR") will return ["pt_BR", "pt", "en"] since these
// locales in that order will be used to find a faker definition
func getFallbackLocales(locale string) []string {
	var result []string
	used := make(map[string]bool)
	result = append(result, locale)
	used[locale] = true
	if strings.Contains(locale, "-") {
		prefix := strings.Split(locale, "-")[0]
		if existsLocale(prefix) {
			result = append(result, prefix)
			used[prefix] = true
		}
	}

	const defaultLocale = "en"
	if !used[defaultLocale] {
		result = append(result, defaultLocale)
	}

	return result
}

// traverse returns the leaf value in the data tree for a given composite key.
// For example: data.traverse("address.street") returns something like data["address"]["street"].
func (d *localeData) traverse(key string) interface{} {
	keys := keys(key)
	var current interface{} = d.data
	for _, k := range keys {
		if m, ok := current.(map[interface{}]interface{}); ok {
			current = m[k]
		} else {
			return nil
		}
	}
	return current
}

func createStringArray(a []interface{}) []string {
	b := make([]string, len(a))
	for i := range a {
		b[i] = fmt.Sprint(a[i])
	}
	return b
}

// Get returns a fake data definition for a given key.
// A fake data definition is a list from which a random value may be
// chosen.
func (d *localeData) Get(key string) []string {
	current := d.traverse(key)
	if _, ok := current.(map[interface{}]interface{}); ok {
		return nil
	} else if a, ok := current.([]interface{}); ok {
		return createStringArray(a)
	} else if current != nil {
		return []string{fmt.Sprint(current)}
	}
	return nil
}

// GetList returns a list of lists of valid values for a given key. This is used for "bullshit" like sentences
// which are built of a valid first, second and third word (subject, verb and object).
// An example for a "bullshit" sentence is "implement value-added web-readyness".
func (d *localeData) GetList(key string) [][]string {
	current := d.traverse(key)
	if a, ok := current.([]interface{}); ok {
		if len(a) > 0 {
			b := a[0]
			if _, ook := b.([]interface{}); ook {
				c := make([][]string, len(a))
				for i := range a {
					b := a[i]
					if d, ok := b.([]interface{}); ok {
						c[i] = createStringArray(d)
					} else {
						return nil
					}
				}
				return c
			}
		}
	}
	return nil
}

func createStringMap(m map[interface{}]interface{}) map[string]string {
	result := make(map[string]string)

	for k, v := range m {
		result[fmt.Sprint(k)] = fmt.Sprint(v)
	}

	return result
}

func createMapArray(a []interface{}) []map[string]string {
	var result []map[string]string

	for _, x := range a {
		if m, ok := x.(map[interface{}]interface{}); ok {
			result = append(result, createStringMap(m))
		}
	}

	return result
}

func (d *localeData) GetStringMapList(key string) []map[string]string {
	current := d.traverse(key)
	if a, ok := current.([]interface{}); ok {
		if len(a) > 0 {
			return createMapArray(a)
		}
	}
	return nil
}

// loadLocaleData loads the fake data definition for a given locale.
func loadLocaleData(locale string) (interface{}, error) {
	bla, err := data.Asset("data/" + locale + ".yml")
	if err != nil {
		return nil, err
	}
	yamlData := make(map[interface{}]interface{})
	err = yaml.Unmarshal(bla, &yamlData)
	if err != nil {
		return nil, err
	}
	m := yamlData[locale]
	return m.(map[interface{}]interface{})["faker"], nil
}

// NewData creates a new fake data definition object for a given locale.
func newData(locale string) (fakerData, error) {
	if r, ok := allDatas[locale]; ok {
		return r, nil
	}

	if !existsLocale(locale) {
		return nil, fmt.Errorf("invalid locale %s", locale)
	}

	locales := getFallbackLocales(locale)
	var data []*localeData
	for _, l := range locales {
		yamlData, err := loadLocaleData(l)
		if err != nil {
			return nil, err
		}
		data = append(data, &localeData{l, yamlData})
	}

	r := &allData{data}
	allDatas[locale] = r
	return r, nil
}

// Get returns a fake data definition for a given key.
// A fake data definition is a list from which a random value may be
// chosen.
func (d *allData) Get(key string) []string {
	for _, locale := range d.locales {
		a := locale.Get(key)
		if a != nil {
			return a
		}
	}
	return nil
}

// GetList returns a list of lists of valid values for a given key. This is used for "bullshit" like sentences
// which are built of a valid first, second and third word (subject, verb and object).
// An example for a "bullshit" sentence is "implement value-added web-readyness".
func (d *allData) GetList(key string) [][]string {
	for _, locale := range d.locales {
		a := locale.GetList(key)
		if a != nil {
			return a
		}
	}
	return nil
}

func (d *allData) GetStringMapList(key string) []map[string]string {
	for _, locale := range d.locales {
		a := locale.GetStringMapList(key)
		if a != nil {
			return a
		}
	}
	return nil
}

func init() {
	allLocales = loadAllLocales()
}
