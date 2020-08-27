package validators

import (
	"regexp"
	"strings"
)

var locatorRegex *regexp.Regexp

func init() {
	locatorRegex = regexp.MustCompile(`^[A-R]{2}[0-9]{2}[A-X]{2}$`)
}

// Locator returns true if locator is valid 6 characters Maidenhead locator.
// locator is case insensitive
func Locator(locator string) bool {
	return locatorRegex.MatchString(strings.ToUpper(locator))
}
