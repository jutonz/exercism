package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0

  for _, v := range birdsPerDay {
    total += v
  }

	return total
}

const daysInWeek = 7
// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	total := 0
	weekOffset := (week - 1) * daysInWeek

	for i := 0; i < daysInWeek; i++ {
		total += birdsPerDay[i + weekOffset]
	}

	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
  for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
  }
  
	return birdsPerDay
}
