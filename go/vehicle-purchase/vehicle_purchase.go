package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	var car = "car"
	var truck = "truck"

	return (kind == car || kind == truck)

}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	const betterChoice = " is clearly the better choice."
	if option1 < option2 {
		return option1 + betterChoice
	} else {
		return option2 + betterChoice
	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	//age < 3 => 80% original price
	//age >10 => 50% originale price
	//age > 3 && <10 => 70%
	var dealingPrice float64 = 0.0

	if age < 3 {
		dealingPrice = originalPrice * 0.80
	} else if age >= 10 {
		dealingPrice = originalPrice / 2
	} else {
		dealingPrice = originalPrice * 0.70
	}

	return dealingPrice
}
