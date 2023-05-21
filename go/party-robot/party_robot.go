package partyrobot

import (
	"fmt"
	"strings"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %v!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %v! You are now %v years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	welcome := Welcome(name)
	assignedTo := fmt.Sprintf("You have been assigned to table %03d. Your table is %v, exactly %.1f meters from here.", table, direction, distance)
	nextTo := fmt.Sprintf("You will be sitting next to %v.", neighbor)

	lines := []string{welcome, assignedTo, nextTo}
	return strings.Join(lines, "\n")
}
