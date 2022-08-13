package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	} else {
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		vehicleDec := option1 + " is clearly the better choice."
		return vehicleDec
	} else {
		vehicleDec := option2 + " is clearly the better choice."
		return vehicleDec
	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		resellPrice := originalPrice * 0.8
		return resellPrice
	} else if age >= 3 && age < 10 {
		resellPrice := originalPrice * 0.7
		return resellPrice
	} else {
		resellPrice := originalPrice * 0.5
		return resellPrice
	}
}
