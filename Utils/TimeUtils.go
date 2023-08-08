package Utils

import "time"

// GetJuly272023Beginning
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

// GetTodayBeginning
// Gets the beginning of the current day
func GetTodayBeginning() time.Time {
	today := time.Now()
	year, month, day := today.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, today.Location())
}

// GetTodayEnd
// Returns the end of the current day
func GetTodayEnd() time.Time {
	today := time.Now()
	year, month, day := today.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, today.Location())
}
