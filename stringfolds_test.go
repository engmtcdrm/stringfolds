package stringfolds

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for [CutPrefixFold] function.
func Test_CutPrefixFold(t *testing.T) {
	tests := []struct {
		s      string
		prefix string
		want   string
		found  bool
	}{
		{"Hello, World!", "hello", ", World!", true},
		{"Hello, World!", "HELLO", ", World!", true},
		{"Hello, World!", "world", "Hello, World!", false},
		{"Hello, World!", "WORLD", "Hello, World!", false},
		{"Hello, World!", "", "Hello, World!", true},
		{"Hello, World!", "Hello, World! But LONGER", "Hello, World!", false},
	}

	for _, tt := range tests {
		if got, found := CutPrefixFold(tt.s, tt.prefix); got != tt.want || found != tt.found {
			require.Equal(t, tt.want, got, "CutPrefixFold(%q, %q) = %q; want %q", tt.s, tt.prefix, got, tt.want)
			require.Equal(t, tt.found, found, "CutPrefixFold(%q, %q) found = %v; want %v", tt.s, tt.prefix, found, tt.found)
		}
	}
}

// Tests for [CutSuffixFold] function.
func Test_CutSuffixFold(t *testing.T) {
	tests := []struct {
		s      string
		suffix string
		want   string
		found  bool
	}{
		{"Hello, World!", "world!", "Hello, ", true},
		{"Hello, World!", "WORLD!", "Hello, ", true},
		{"Hello, World!", "hello", "Hello, World!", false},
		{"Hello, World!", "HELLO", "Hello, World!", false},
		{"Hello, World!", "", "Hello, World!", true},
		{"Hello, World!", "Hello, World! But LONGER", "Hello, World!", false},
	}

	for _, tt := range tests {
		if got, found := CutSuffixFold(tt.s, tt.suffix); got != tt.want || found != tt.found {
			require.Equal(t, tt.want, got, "CutSuffixFold(%q, %q) = %q; want %q", tt.s, tt.suffix, got, tt.want)
			require.Equal(t, tt.found, found, "CutSuffixFold(%q, %q) found = %v; want %v", tt.s, tt.suffix, found, tt.found)
		}
	}
}

// Tests for [HasPrefixFold] function.
func Test_HasPrefixFold(t *testing.T) {
	tests := []struct {
		s      string
		prefix string
		want   bool
	}{
		{"Hello, World!", "hello", true},
		{"Hello, World!", "HELLO", true},
		{"Hello, World!", "world", false},
		{"Hello, World!", "WORLD", false},
		{"Hello, World!", "", true},
		{"Hello, World!", "Hello, World! But LONGER", false},
	}

	for _, tt := range tests {
		if got := HasPrefixFold(tt.s, tt.prefix); got != tt.want {
			require.Equal(t, tt.want, got, "HasPrefixFold(%q, %q) = %v; want %v", tt.s, tt.prefix, got, tt.want)
		}
	}
}

// Tests for [HasSuffixFold] function.
func Test_HasSuffixFold(t *testing.T) {
	tests := []struct {
		s      string
		suffix string
		want   bool
	}{
		{"Hello, World!", "world!", true},
		{"Hello, World!", "WORLD!", true},
		{"Hello, World!", "hello", false},
		{"Hello, World!", "HELLO", false},
		{"Hello, World!", "", true},
		{"Hello, World!", "Hello, World! But LONGER", false},
	}

	for _, tt := range tests {
		if got := HasSuffixFold(tt.s, tt.suffix); got != tt.want {
			require.Equal(t, tt.want, got, "HasSuffixFold(%q, %q) = %v; want %v", tt.s, tt.suffix, got, tt.want)
		}
	}
}

// Tests for [TrimPrefixFold] function.
func Test_TrimPrefixFold(t *testing.T) {
	tests := []struct {
		s      string
		prefix string
		want   string
	}{
		{"Hello, World!", "hello", ", World!"},
		{"Hello, World!", "HELLO", ", World!"},
		{"Hello, World!", "world", "Hello, World!"},
		{"Hello, World!", "WORLD", "Hello, World!"},
		{"Hello, World!", "", "Hello, World!"},
		{"Hello, World!", "Hello, World! But LONGER", "Hello, World!"},
	}

	for _, tt := range tests {
		if got := TrimPrefixFold(tt.s, tt.prefix); got != tt.want {
			require.Equal(t, tt.want, got, "TrimPrefixFold(%q, %q) = %q; want %q", tt.s, tt.prefix, got, tt.want)
		}
	}
}

// Tests for [TrimSuffixFold] function.
func Test_TrimSuffixFold(t *testing.T) {
	tests := []struct {
		s      string
		suffix string
		want   string
	}{
		{"Hello, World!", "world!", "Hello, "},
		{"Hello, World!", "WORLD!", "Hello, "},
		{"Hello, World!", "hello", "Hello, World!"},
		{"Hello, World!", "HELLO", "Hello, World!"},
		{"Hello, World!", "", "Hello, World!"},
		{"Hello, World!", "Hello, World! But LONGER", "Hello, World!"},
	}

	for _, tt := range tests {
		if got := TrimSuffixFold(tt.s, tt.suffix); got != tt.want {
			require.Equal(t, tt.want, got, "TrimSuffixFold(%q, %q) = %q; want %q", tt.s, tt.suffix, got, tt.want)
		}
	}
}
