package common

import (
	"regexp"
	"unicode/utf8"
)

const slugValidPattern = `^[A-Z_0-9]+$`

func ValidateSlug(slug string) (bool, string, error) {
	if utf8.RuneCountInString(slug) > 256 {
		return false, "slug length exceeded", nil
	}

	regExp, err := regexp.Compile(slugValidPattern)
	if err != nil {
		return false, "invalid regular expression", ErrCompileRegexpr
	}

	if !regExp.MatchString(slug) {
		return false, "invalid slug format", nil
	}

	return true, "", nil
}
