package lib

var Multiple float64 = 100000

func MoneyIn(in float64) float64 {
	return Multiple * in
}

func MoneyOut(in float64) float64 {
	return in / Multiple
}
