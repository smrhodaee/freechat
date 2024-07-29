package validations

import "strings"

func IsValidGroupName(name string) bool {
	if len(name) < 3 {
		return false
	}
	if strings.Contains(name, ".") {
		return false
	}
	return true
}

func IsValidRoomTitle(title string) bool {
	return len(title) > 3
}
