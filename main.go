package main

import (
	"github.com/romankravchuk/learn-go/book"
)

// Start program point
func main() {
	// showGopher()
	var publisher, writer, artist, title, ganre string
	var year, pageNumber int
	var grade, cost float32

	title, writer = "Mr. GoToSleep", "Tracey Hatchet"
	artist, publisher = "Jewel Tampson", "DizzyBooks Publishing Inc."
	ganre = "Romance"
	year, pageNumber = 1997, 14
	grade = 6.5
	cost = book.CalcCost(year, pageNumber, grade)

	book.ShowBookInfo(publisher, writer, artist, title, pageNumber, year, grade, ganre, cost)

	title, writer = "Epic Vol. 1", "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	ganre = "Si-Fi"
	year, pageNumber = 2013, 160
	grade = 9.0
	cost = book.CalcCost(year, pageNumber, grade)

	book.ShowBookInfo(publisher, writer, artist, title, pageNumber, year, grade, ganre, cost)
}
