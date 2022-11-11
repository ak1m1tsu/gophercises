package main

import (
	"fmt"

	"github.com/romankravchuk/learn-go/lib/book"
)

func main() {
	title, writer := "Mr. GoToSleep", "Tracey Hatchet"
	artist, publisher := "Jewel Tampson", "DizzyBooks Publishing Inc."
	ganre := "Romance"
	year, pageNumber := 1997, 14
	grade := 6.5
	cost := book.CalcCost(year, pageNumber, grade)

	fmt.Println(title, "written by", writer, "drawn by", artist)
	fmt.Println("Pages:", pageNumber, "\nGrade:", grade, "\nGanre:", ganre)
	fmt.Println("Published by", publisher, "in", year)
	fmt.Println("Cost:", cost)
	fmt.Println()
}
