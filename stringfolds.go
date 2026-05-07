package stringfolds

import (
	"strings"
)

// HasPrefixFold is a case-insensitive version of [strings.HasPrefix].
func HasPrefixFold(s, prefix string) bool {
	n := len(prefix)
	if len(s) < n {
		return false
	}

	return strings.EqualFold(s[:n], prefix)
}

// HasSuffixFold is a case-insensitive version of [strings.HasSuffix].
func HasSuffixFold(s, suffix string) bool {
	n := len(suffix)
	if len(s) < n {
		return false
	}

	return strings.EqualFold(s[len(s)-n:], suffix)
}

// TrimPrefixFold is a case-insensitive version of [strings.TrimPrefix].
func TrimPrefixFold(s, prefix string) string {
	n := len(prefix)
	if len(s) < n {
		return s
	}

	if strings.EqualFold(s[:n], prefix) {
		return s[n:]
	}

	return s
}

// TrimSuffixFold is a case-insensitive version of [strings.TrimSuffix].
func TrimSuffixFold(s, suffix string) string {
	n := len(suffix)
	if len(s) < n {
		return s
	}

	if strings.EqualFold(s[len(s)-n:], suffix) {
		return s[:len(s)-n]
	}

	return s
}
