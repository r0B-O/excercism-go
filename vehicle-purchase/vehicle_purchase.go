package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		vehicleDec := option1 + " is clearly the better choice."
		return vehicleDec
	}
	vehicleDec := option2 + " is clearly the better choice."
	return vehicleDec
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	percent := 0.5
	if age < 3 {
		percent = 0.8
	} else if age < 10 {
		percent = 0.7
	}
	return originalPrice * percent
}
