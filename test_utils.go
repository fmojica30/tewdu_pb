package main

import "time"

// GetJul272023Beginning
// Gets the beginning of July 27, 2023, for testing purposes
func GetJuly272023Beginning() time.Time {
	today := time.Now()
	return time.Date(2023, 07, 27, 0, 0, 0, 0, today.Location())
}

// GetJuly272023End
// Get the end of July 27, 2023, for testing purposes
func GetJuly272023End() time.Time {
	today := time.Now()
	return time.Date(2023, 07, 27, 23, 59, 59, 0, today.Location())
}
