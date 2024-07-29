package validations

func IsValidMessageText(text string) bool {
	return len(text) > 2
}
