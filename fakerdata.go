package gofaker

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ralli/gofaker/data"
	yaml "gopkg.in/yaml.v2"
)

// Data provides access to fake data definitions used to make up fake data.
type Data interface {
	Get(key string) []string
}

type localeData struct {
	locale string
	data   interface{}
}

// data holds Data onjects for multiple locales.
// The Get operation will try all these locale
// data objects until a match is found.
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
		result = append(result, strings.TrimSuffix(v, ".yml"))
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

// Get returns a fake data definition for a given key.
// A fake data definition is a list from which a random value may be
// chosen.
func (d *localeData) Get(key string) []string {
	keys := keys(key)
	var current interface{} = d.data
	for _, k := range keys {
		if m, ok := current.(map[interface{}]interface{}); ok {
			current = m[k]
		} else {
			return nil
		}
	}
	if _, ok := current.(map[interface{}]interface{}); ok {
		return nil
	} else if a, ok := current.([]interface{}); ok {
		b := make([]string, len(a))
		for i := range a {
			b[i] = fmt.Sprint(a[i])
		}
		return b
	} else if current != nil {
		return []string{fmt.Sprint(current)}
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
func NewData(locale string) (Data, error) {
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

func init() {
	allLocales = loadAllLocales()
}
