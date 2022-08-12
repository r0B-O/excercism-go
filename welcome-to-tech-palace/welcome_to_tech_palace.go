package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	welMsg := "Welcome to the Tech Palace, " + strings.ToUpper(customer)

	return welMsg
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	repStar := strings.Repeat("*", numStarsPerLine)
	welMsg := repStar + "\n" + welcomeMsg + "\n" + repStar
	return welMsg
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	starRemoved := strings.ReplaceAll(oldMsg, "*", "")
	trimmedMsg := strings.TrimSpace(starRemoved)
	return trimmedMsg
}
