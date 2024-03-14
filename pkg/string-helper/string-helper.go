package stringhelper

func TruncateText(t string, maxLength int) string {
	if len(t) <= maxLength {
		return t
	}
	return t[:maxLength] + "..."
}
