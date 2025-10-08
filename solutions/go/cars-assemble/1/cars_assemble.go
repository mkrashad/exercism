package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    var workingCarsPerHour = float64(productionRate) * successRate / 100
	return int(workingCarsPerHour / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var groupsOfTen = uint(carsCount) / 10
	var groupsIndividual = uint(carsCount) % 10 
    return groupsOfTen * 95000 + groupsIndividual * 10000
}
