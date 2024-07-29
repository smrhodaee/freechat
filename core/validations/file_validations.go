package validations

func IsValidFilename(filename string) bool {
	if len(filename) > 100 {
		return false
	}
	return true 
}