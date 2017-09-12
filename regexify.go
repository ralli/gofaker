package gofaker

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Regexify generates a string that matches a given regular expression.
//
// Given a regular expression, attempt to generate a string
// that would match it.  This is a rather simple implementation,
// so don't be shocked if it blows up on you in a spectacular fashion.
//
// It does not handle ., *, unbounded ranges such as {1,},
// extensions such as (?=), character classes, some abbreviations
// for character classes, and nested parentheses.
//
// I told you it was simple. :) It's also probably pretty slow,
// so you shouldn't use it.
//
// It will take a regex like this:
//
// /^[A-PR-UWYZ0-9][A-HK-Y0-9][AEHMNPRTVXY0-9]?[ABEHMNPRVWXY0-9]? {1,2}[0-9][ABD-HJLN-UW-Z]{2}$/
//
// and generate a string like this:
//
// "U3V  3TP"
//
func (f *Faker) Regexify(expression string) string {
	// remove optional leading slash and caret
	s := gsub("^\\/?\\^?", expression, "")
	// remove optional trailing slash and dollar
	s = gsub("\\$?\\/?$", s, "")
	// All {2} become {2,2} and ? become {0,1}
	s = gsub("\\{(\\d+)\\}", s, "{$1,$1}")
	// All ? become {0,1}
	s = gsub("\\?", s, "{0,1}")
	// [12]{1,2} becomes [12] or [12][12]
	s = f.replaceGroups("(\\[[^\\]]+\\])\\{(\\d+),(\\d+)\\}", s)
	// (12|34){1,2} becomes (12|34) or (12|34)(12|34)
	s = f.replaceGroups("(\\([^\\)]+\\))\\{(\\d+),(\\d+)\\}", s)
	// A{1,2} becomes A or AA or \d{3} becomes \d\d\d
	s = f.replaceGroups("(\\\\.)\\{(\\d+),(\\d+)\\}", s)
	s = f.replaceGroups("(.)\\{(\\d+),(\\d+)\\}", s)
	s = gsubFunc("\\[([^\\]]+)\\]", s, func(aaa string) string {
		return gsubFunc("\\w-\\w", aaa, func(str string) string {
			r1 := str[0]
			r2 := str[2]
			n := rune(f.random.Intn(int(r2-r1+1)) + int(r1))
			return fmt.Sprintf("%c", n)
		})
	})

	s = gsubFunc("\\[([^\\]]+)\\]", s, func(aaa string) string {
		bla := gsub("(\\[|\\])", aaa, "")
		return fmt.Sprintf("%c", bla[f.random.Intn(len(bla))])
	})

	s = gsubFunc("\\([^\\)]+\\)", s, func(aaa string) string {
		parts := strings.Split(gsub("(\\(|\\))", aaa, ""), "|")
		return f.randomValue(parts)
	})

	s = gsubFunc("\\\\d", s, func(aaa string) string {
		return fmt.Sprintf("%c", f.RandomDigit())
	})

	s = gsubFunc("\\\\w", s, func(aaa string) string {
		return fmt.Sprintf("%c", f.RandomLetter())
	})

	return s
}

func gsubFunc(exp, expression string, repl func(string) string) string {
	re, err := regexp.Compile(exp)
	if err != nil {
		panic(err)
	}
	return re.ReplaceAllStringFunc(expression, repl)
}

func gsub(expression, input, replacement string) string {
	re, err := regexp.Compile(expression)
	if err != nil {
		panic(err)
	}
	return re.ReplaceAllString(input, replacement)
}

func (f *Faker) replaceGroup(exp, expression string) (string, bool) {
	re, err := regexp.Compile(exp)
	if err != nil {
		panic(err)
	}
	submatches := re.FindStringSubmatch(expression)
	if submatches == nil {
		return expression, false
	}
	from, err := strconv.Atoi(submatches[2])
	if err != nil {
		panic(err)
	}
	to, err := strconv.Atoi(submatches[3])
	if err != nil {
		panic(err)
	}
	group := submatches[1]
	n := f.random.Intn(to-from+1) + from
	replacement := strings.Repeat(group, n)
	result := strings.Replace(expression, submatches[0], replacement, 1)
	return result, true
}

func (f *Faker) replaceGroups(exp, expression string) string {
	result := expression
	var ok bool
	for i := 0; i < 10; i++ {
		result, ok = f.replaceGroup(exp, result)
		if !ok {
			return result
		}
	}
	return result
}
