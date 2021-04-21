package common

import (
	"regexp"
	"strconv"
	"time"
)

func RemoveSpaces(input string) string {
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(input, " ")
}

func CheckNumber(input string) bool {
	_, err := strconv.ParseUint(input, 10, 0)
	if err == nil {
		return true
	} else {
		return false
	}
}

func CheckDateString(input string) bool {
	_, err := time.Parse("2006-01-02", input)
	if err == nil {
		return true
	} else {
		return false
	}
}
