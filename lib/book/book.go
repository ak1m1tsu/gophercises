package book

// Calculate book price
func CalcCost(year int, pageNumber int, grade float64) float64 {
	return float64(2022-year) / 100.00 * float64(pageNumber*100) * grade
}
