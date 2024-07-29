package tools

func FindElement[T any](lst []T, isValid func(*T) bool) *T {
	for _, el := range lst {
		if isValid(&el) {
			return &el
		}
	}
	return nil
}
