package parsing

import (
	"strconv"
	"strings"
)

// ConcatIfNotEmpty will concatenate the given strings with the given separator
// only if the string is not equal to ""
func ConcatIfNotEmpty(sep string, parts ...string) string {
	var nonempty []string
	for _, s := range parts {
		if s != "" {
			nonempty = append(nonempty, s)
		}
	}
	return strings.Join(nonempty, sep)
}

// ParseIntOrDefault will attempt to parse the given string as an integer and
// if an error is generated, then returned the supplied default instead
func ParseIntOrDefault(s string, d int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return d
	}
	return i
}

// ParseInt64OrDefault will attempt to parse the given string as an integer and
// if an error is generated, then returned the supplied default instead
func ParseInt64OrDefault(s string, d int64) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return d
	}
	return i
}

// ParseFloatOrDefault will attempt to parse the given string as a float and
// if an error is generated, then returned the supplied default instead
func ParseFloatOrDefault(s string, d float64) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return d
	}
	return f
}

// ParseYesOrNo will return true if the given string equals "yes" in a case-
// insensitive manner, otherwise it will return false
func ParseYesOrNo(s string) bool {
	return strings.ToLower(s) == "yes"
}
