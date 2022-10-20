package main

import (
	bankheist "github.com/romankravchuk/learn-go/bank-heist"
	"github.com/romankravchuk/learn-go/book"
	"github.com/romankravchuk/learn-go/gopher"
)

// Start program point
func main() {
	gopher.PrintGopher()
	book.RunBookApp()
	bankheist.RunBankHeist()
}
