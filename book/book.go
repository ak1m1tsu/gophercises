package book

import "fmt"

func RunBookApp() {
	title, writer := "Mr. GoToSleep", "Tracey Hatchet"
	artist, publisher := "Jewel Tampson", "DizzyBooks Publishing Inc."
	ganre := "Romance"
	year, pageNumber := 1997, 14
	grade := 6.5
	cost := calcCost(year, pageNumber, grade)

	showBookInfo(publisher, writer, artist, title, ganre, pageNumber, year, grade, cost)

	title, writer = "Epic Vol. 1", "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	ganre = "Si-Fi"
	year, pageNumber = 2013, 160
	grade = 9.0
	cost = calcCost(year, pageNumber, grade)

	showBookInfo(publisher, writer, artist, title, ganre, pageNumber, year, grade, cost)
}

// Calculate book price
func calcCost(year int, pageNumber int, grade float64) float64 {
	return float64(2022-year) / 100.00 * float64(pageNumber*100) * grade
}

// Print book info
func showBookInfo(
	publisher,
	writer,
	artist,
	title,
	ganre string,
	pageNumber,
	year int,
	grade,
	cost float64) {
	fmt.Println(title, "written by", writer, "drawn by", artist)
	fmt.Println("Pages:", pageNumber, "\nGrade:", grade, "\nGanre:", ganre)
	fmt.Println("Published by", publisher, "in", year)
	fmt.Println("Cost:", cost)
	fmt.Println()
}
