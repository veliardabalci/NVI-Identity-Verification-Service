package utils

import "regexp"

func IsNumeric(input string) bool {
	re := regexp.MustCompile(`^\d+$`)
	return re.MatchString(input)
}

func IsYearValid(year string) bool {
	re := regexp.MustCompile(`^\d{4}$`)
	return re.MatchString(year)
}
