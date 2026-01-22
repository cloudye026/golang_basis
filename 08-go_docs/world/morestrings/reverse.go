package morestrings

// IsString checks if the given value is a string type.
// Returns "是字符串" if true, "不是字符串" if false.
func IsString(s any) string {
	if _, ok := s.(string); ok {
		return "是字符串"
	}
	return "不是字符串"
}
