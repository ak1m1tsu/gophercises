package book

import "fmt"

// Calculate book price
func CalcCost(year int, pageNumber int, grade float32) float32 {
	return float32(2022-year) / 100.00 * float32(pageNumber*100) * grade
}

// Print book info
func ShowBookInfo(
	publisher string,
	writer string,
	artist string,
	title string,
	pageNumber int,
	year int,
	grade float32,
	ganre string,
	cost float32) {
	fmt.Println(title, "written by", writer, "drawn by", artist)
	fmt.Println("Pages:", pageNumber, "\nGrade:", grade, "\nGanre:", ganre)
	fmt.Println("Published by", publisher, "in", year)
	fmt.Println("Cost:", cost)
	fmt.Println()
}
