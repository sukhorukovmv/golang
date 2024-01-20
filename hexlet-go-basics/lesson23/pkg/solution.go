package pkg

import (
	"strings"
	"unicode"
)

// BEGIN (write your solution here)
func LatinLetters(s string) string {
	latinString := strings.Builder{}
	for _, ch := range s {
		if unicode.Is(unicode.Latin, ch) {
			latinString.WriteRune(ch)
		}
	}
	return latinString.String()
}

// END
