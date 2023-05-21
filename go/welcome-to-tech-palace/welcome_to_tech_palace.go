package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starLine := strings.Repeat("*", numStarsPerLine)
	lines := []string{starLine, welcomeMsg, starLine}
	return strings.Join(lines, "\n")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	withoutNewlines := strings.ReplaceAll(oldMsg, "\n", "")
	withoutStars := strings.Trim(withoutNewlines, "*")
	withoutWhitespace := strings.Trim(withoutStars, " ")
	return withoutWhitespace
}
