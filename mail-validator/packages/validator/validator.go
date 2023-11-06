package validator

import (
	"regexp"
)

func ValidateMailString(mail string) bool {
	pattern := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	return pattern.MatchString(mail)
}
