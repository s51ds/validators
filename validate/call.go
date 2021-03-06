package validate

import (
	"strconv"
	"strings"
)

// CallSign returns true if callSign format is valid
func CallSign(callSign string) bool {
	callSignLen := len(callSign)
	// check if len is at least 3 characters
	if callSignLen < 3 {
		return false
	}
	callSign = strings.ToUpper(callSign)
	digits, chars, slashes := 0, 0, 0
	for _, v := range callSign {
		switch {
		case v == 47:
			{ // slash /
				slashes++
			}
		case v >= 48 && v <= 57: // 0-9
			{
				digits++
			}
		case v >= 65 && v <= 90: // A-Z
			{
				chars++
			}
		default:
			return false
		}
	}
	if digits == 0 || chars == 0 || slashes > 2 {
		return false
	}

	if slashes > 0 {
		for _, v := range strings.Split(callSign, "/") {
			if len(v) == 0 {
				return false
			}
		}
	} else {
		// final checks:
		// last character can not be a number
		b := byte(callSign[callSignLen-1])
		if b >= 48 && b <= 57 {
			return false
		}
		// only one number is allowed as first character
		if _, err := strconv.Atoi(callSign[:2]); err == nil {
			return false
		}
	}
	return true
}
