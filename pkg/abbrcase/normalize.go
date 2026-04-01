package abbrcase

import (
	"strings"
	"unicode"
)

// SuggestedName rewrites all-uppercase runs into camel-case abbreviation form.
func SuggestedName(name string) (string, bool) {
	runes := []rune(name)
	if len(runes) == 0 {
		return name, false
	}

	var builder strings.Builder
	builder.Grow(len(name))

	changed := false

	for i := 0; i < len(runes); {
		if !unicode.IsUpper(runes[i]) {
			builder.WriteRune(runes[i])
			i++
			continue
		}

		j := i + 1
		for j < len(runes) && unicode.IsUpper(runes[j]) {
			j++
		}

		end := j
		if j-i >= 2 && j < len(runes) && unicode.IsLower(runes[j]) {
			end = j - 1
		}

		if end-i >= 2 {
			builder.WriteRune(runes[i])
			for _, r := range runes[i+1 : end] {
				builder.WriteRune(unicode.ToLower(r))
			}
			changed = true
			i = end
		} else {
			builder.WriteRune(runes[i])
			i++
		}
	}

	suggestion := builder.String()
	return suggestion, changed && suggestion != name
}
