package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate/60, successRate))
}

const costOf10Cars uint = 95_000
const costOfOneCar uint = 10_000

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupsOf10 := uint(carsCount / 10)
	leftover := uint(carsCount % 10)

	return groupsOf10*costOf10Cars + leftover*costOfOneCar
}
