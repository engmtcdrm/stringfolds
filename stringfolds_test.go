package stringfolds

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
	}

	for _, tt := range tests {
		if got := TrimSuffixFold(tt.s, tt.suffix); got != tt.want {
			require.Equal(t, tt.want, got, "TrimSuffixFold(%q, %q) = %q; want %q", tt.s, tt.suffix, got, tt.want)
		}
	}
}
