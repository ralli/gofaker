package gofaker

import "bytes"

var romanizeTable = map[rune]string{
	// romanize_cyrillic, Based on conventions abopted by BGN/PCGN for Ukrainian
	'а': "a", 'б': "b", 'в': "v", 'г': "h", 'ґ': "g", 'д': "d",
	'е': "e", 'є': "ye", 'ж': "zh", 'з': "z", 'и': "y", 'і': "i",
	'ї': "yi", 'й': "y", 'к': "k", 'л': "l", 'м': "m", 'н': "n",
	'о': "o", 'п': "p", 'р': "r", 'с': "s", 'т': "t", 'у': "u",
	'ф': "f", 'х': "kh", 'ц': "ts", 'ч': "ch", 'ш': "sh", 'щ': "shch",
	'ю': "yu", 'я': "ya",
	'А': "a", 'Б': "b", 'В': "v", 'Г': "h", 'Ґ': "g", 'Д': "d",
	'Е': "e", 'Є': "ye", 'Ж': "zh", 'З': "z", 'И': "y", 'І': "i",
	'Ї': "yi", 'Й': "y", 'К': "k", 'Л': "l", 'М': "m", 'Н': "n",
	'О': "o", 'П': "p", 'Р': "r", 'С': "s", 'Т': "t", 'У': "u",
	'Ф': "f", 'Х': "kh", 'Ц': "ts", 'Ч': "ch", 'Ш': "sh", 'Щ': "shch",
	'Ю': "yu", 'Я': "ya",
	'ь': "", // Ignore symbol, because its standard presentation is not allowed in URLs
	// german umlauts
	'ä': "ae", 'ö': "oe", 'ü': "ue", 'Ä': "Ae", 'Ö': "Oe", 'Ü': "Ue", 'ß': "ss",
}

// Romanize maps characters from a s to the ascii range. Currently supports
// cyrillic characters and german umlauts.
func Romanize(s string) string {
	var out bytes.Buffer
	for _, c := range s {
		if val, ok := romanizeTable[c]; ok {
			out.Write([]byte(val))
		} else {
			out.WriteRune(c)
		}
	}
	return string(out.Bytes())
}
