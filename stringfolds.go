package stringfolds

import (
	"strings"
)

// CutPrefixFold is a case-insensitive version of [strings.CutPrefix].
func CutPrefixFold(s, prefix string) (after string, found bool) {
	if prefix == "" {
		return s, true
	}

	if HasPrefixFold(s, prefix) {
		return s[len(prefix):], true
	}

	return s, false
}

// CutSuffixFold is a case-insensitive version of [strings.CutSuffix].
func CutSuffixFold(s, suffix string) (before string, found bool) {
	if suffix == "" {
		return s, true
	}

	if HasSuffixFold(s, suffix) {
		return s[:len(s)-len(suffix)], true
	}

	return s, false
}

// HasPrefixFold is a case-insensitive version of [strings.HasPrefix].
func HasPrefixFold(s, prefix string) bool {
	if len(prefix) == 0 {
		return true
	}

	n := len(prefix)
	if len(s) < n {
		return false
	}

	return strings.EqualFold(s[:n], prefix)
}

// HasSuffixFold is a case-insensitive version of [strings.HasSuffix].
func HasSuffixFold(s, suffix string) bool {
	if len(suffix) == 0 {
		return true
	}

	n := len(suffix)
	if len(s) < n {
		return false
	}

	return strings.EqualFold(s[len(s)-n:], suffix)
}

// TrimPrefixFold is a case-insensitive version of [strings.TrimPrefix].
func TrimPrefixFold(s, prefix string) string {
	if prefix == "" {
		return s
	}

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
	if suffix == "" {
		return s
	}

	n := len(suffix)
	if len(s) < n {
		return s
	}

	if strings.EqualFold(s[len(s)-n:], suffix) {
		return s[:len(s)-n]
	}

	return s
}
