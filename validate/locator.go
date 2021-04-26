package validate

import (
	"regexp"
)

var locatorRegex6chars *regexp.Regexp
var locatorRegex4chars *regexp.Regexp

func init() {
	locatorRegex6chars = regexp.MustCompile(`^[A-Ra-r]{2}[0-9]{2}[A-Xa-x]{2}$`)
	locatorRegex4chars = regexp.MustCompile(`^[A-Ra-r]{2}[0-9]{2}$`)
}

// Locator returns true if locator is valid 6 characters Maidenhead locator.
// locator is case insensitive
func Locator(locator string) bool {
	if !locatorRegex6chars.MatchString(locator) {
		if locatorRegex4chars.MatchString(locator) {
			return true
		}
	} else {
		return true
	}
	return false

}
