package strcon

import (
	"bytes"
	"unicode"
)

// Swap character case from upper to lower OR lower to upper
func Swapcase(str string) string {
	buf := &bytes.Buffer{}
	for _, r := range str {
		if unicode.IsUpper(r) {
			buf.WriteRune(unicode.ToLower(r))
		} else {
			buf.WriteRune(unicode.ToUpper(r))
		}
	}
	return buf.String()
}
