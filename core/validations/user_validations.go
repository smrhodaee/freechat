package validations

import "regexp"

const (
	maxUsernameLen = 30
	minUsernameLen = 3
)

func IsValidUsername(username string) bool {
	length := len(username)
	matched, _ := regexp.MatchString("^[A-Za-z0-9]+$", username)
	if !matched || length > maxUsernameLen || length < minUsernameLen {
		return false
	}
	return true
}

func IsValidPassword(password string) bool {
	return true
}
