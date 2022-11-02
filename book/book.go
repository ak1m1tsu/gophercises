package book

import "fmt"

func Run() {
	title, writer := "Mr. GoToSleep", "Tracey Hatchet"
	artist, publisher := "Jewel Tampson", "DizzyBooks Publishing Inc."
	ganre := "Romance"
	year, pageNumber := 1997, 14
	grade := 6.5
	cost := calcCost(year, pageNumber, grade)

	fmt.Println(title, "written by", writer, "drawn by", artist)
	fmt.Println("Pages:", pageNumber, "\nGrade:", grade, "\nGanre:", ganre)
	fmt.Println("Published by", publisher, "in", year)
	fmt.Println("Cost:", cost)
	fmt.Println()
}

// Calculate book price
func calcCost(year int, pageNumber int, grade float64) float64 {
	return float64(2022-year) / 100.00 * float64(pageNumber*100) * grade
}
