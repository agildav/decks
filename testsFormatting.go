package main

// Color formatting
const (
	BeginErrorColor = "\033[1;31m"
	EndErrorColor   = "\033[0m"
)

// Color the error message
func colorError(s string) string {
	return BeginErrorColor + s + EndErrorColor
}
