package techpalace

import "strings"

const welcomeMessage string = "Welcome to the Tech Palace, "

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return welcomeMessage + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var borders = strings.Repeat("*", numStarsPerLine)
	return borders + "\n" + welcomeMsg + "\n" + borders
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	var cleanedMsg = strings.ReplaceAll(oldMsg, "*", "")
	cleanedMsg = strings.ReplaceAll(cleanedMsg, "\n", "")
	return strings.Trim(cleanedMsg, " ")
}
