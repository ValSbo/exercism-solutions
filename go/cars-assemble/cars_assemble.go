package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {

	result := float64(productionRate) * (successRate / 100)
	return result
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	//panic("CalculateWorkingCarsPerMinute not implemented")
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const singleCarCost int = 10000
	const tenCarsCost int = 95000

	groupTenCars := carsCount / 10
	singleCars := carsCount % 10

	return uint((groupTenCars * tenCarsCost) + (singleCars * singleCarCost))

}
