package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    a:=float64(productionRate)
    return a * (successRate / 100)
	panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    a:= float64(productionRate)
    b:= successRate  
    result:= a * (b /100)

    final:= result / 60
    return int(final)
	panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    a:=(carsCount / 10) * 95000
    b:=(carsCount % 10) * 10000
    c:=a + b
    return uint(c)
	panic("CalculateCost not implemented")
}
