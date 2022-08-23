package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	welcomeMsg := fmt.Sprintf("Welcome to my party, %s!", name)

	return welcomeMsg
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	bdayWish := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
	return bdayWish
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	greet := Welcome(name)
	tableAssgn := fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.", table, direction, distance)
	seatMate := fmt.Sprintf("You will be sitting next to %s.", neighbor)
	guidanceMsg := fmt.Sprintf("%s\n%s\n%s", greet, tableAssgn, seatMate)

	return guidanceMsg
}
