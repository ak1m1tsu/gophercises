package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"log"
)

const (
	bannerHt = 95.0
	xIndent  = 40.0
	taxRate  = 0.09
)

type LineItem struct {
	UnitName       string
	PricePerUnit   int
	UnitsPurchased int
}

func main() {
	lineItems := []LineItem{
		{
			UnitName:       "2x6 Lumber - 8'",
			PricePerUnit:   375, // in cents
			UnitsPurchased: 220,
		}, {
			UnitName:       "Drywall Sheet",
			PricePerUnit:   822, // in cents
			UnitsPurchased: 50,
		}, {
			UnitName:       "Paint",
			PricePerUnit:   1455, // in cents
			UnitsPurchased: 3,
		},
	}
	subtotal := 0
	for _, li := range lineItems {
		subtotal += li.PricePerUnit * li.UnitsPurchased
	}
	tax := int(float64(subtotal) * taxRate)
	total := subtotal + tax
	totalUSD := toUSD(total)
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	width, height := pdf.GetPageSize()
	log.Printf("Page size: width=%v, height=%v\n", width, height)
	pdf.AddPage()

	pdf.SetFillColor(103, 60, 79)
	pdf.Polygon([]gofpdf.PointType{
		{0, 0},
		{width, 0},
		{width, bannerHt},
		{0, bannerHt * 0.9},
	}, "F")
	pdf.Polygon([]gofpdf.PointType{
		{0, height},
		{0, height - (bannerHt * 0.2)},
		{width, height - (bannerHt * 0.1)},
		{width, height},
	}, "F")

	pdf.SetFont("arial", "B", 40)
	pdf.SetTextColor(255, 255, 255)
	_, lineHt := pdf.GetFontSize()
	pdf.Text(xIndent, bannerHt-bannerHt/2.0+lineHt/3.0, "INVOICE")

	pdf.SetFont("arial", "", 12)
	pdf.SetTextColor(255, 255, 255)
	_, lineHt = pdf.GetFontSize()
	pdf.MoveTo(width-xIndent-1.75*175.0, (bannerHt-lineHt*4.5)/2.0)
	pdf.MultiCell(
		175.0,
		lineHt*1.5,
		"(123) 456-7890\nkravchuk.roman02@gmail.com\nak1m1tsu.app",
		gofpdf.BorderNone,
		gofpdf.AlignRight,
		false,
	)
	pdf.MoveTo(width-xIndent-175.0, (bannerHt-lineHt*4.5)/2.0)
	pdf.MultiCell(
		175.0,
		lineHt*1.5,
		"123 Fake St\nSome Town, PA\n12345",
		gofpdf.BorderNone,
		gofpdf.AlignRight,
		false,
	)

	_, sy := summaryBlock(pdf, xIndent, bannerHt+lineHt*2.0, "Billed to", "Client Name", "Client Address", "City, State, Country", "Postal Code")

	summaryBlock(pdf, xIndent*2.0+lineHt*12.5, bannerHt+lineHt*2.0, "Invoice number", "00000000123")
	summaryBlock(pdf, xIndent*2.0+lineHt*12.5, bannerHt+lineHt*6.75, "Date of Issue", "29.12.2022")

	x, y := width-xIndent-175.0, bannerHt+lineHt*2.25
	pdf.MoveTo(x, y)
	pdf.SetFont("times", "", 14)
	_, lineHt = pdf.GetFontSize()
	pdf.SetTextColor(180, 180, 180)
	pdf.CellFormat(175.0, lineHt, "Invoice Total", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x, y = x+5, y+lineHt*2.5
	pdf.MoveTo(x, y)
	pdf.SetFontSize(48)
	alpha := 58
	pdf.SetTextColor(72+alpha, 42+alpha, 55+alpha)
	pdf.CellFormat(175.0, lineHt, totalUSD, gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x, y = x-5, y+lineHt*1.25

	if sy > y {
		y = sy
	}
	x, y = xIndent-20.0, y+30.0
	pdf.Rect(x, y, width-xIndent*2.0+40.0, 3.0, "F")

	pdf.SetFont("times", "", 14)
	_, lineHt = pdf.GetFontSize()
	pdf.SetTextColor(180, 180, 180)
	x, y = xIndent-2.0, y+lineHt
	pdf.MoveTo(x, y)
	pdf.CellFormat(width/2.65+1.0, lineHt, "Description", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignLeft, false, 0, "")
	x = x + width/2.25 - xIndent
	pdf.MoveTo(x, y)
	pdf.CellFormat(100, lineHt, "Price Per Unit", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = x + 100.0
	pdf.MoveTo(x, y)
	pdf.CellFormat(80, lineHt, "Quantity", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = width - xIndent - 2.0 - 120.0
	pdf.MoveTo(x, y)
	pdf.CellFormat(120, lineHt, "Amount", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")

	y = y + lineHt
	for _, lineItem := range lineItems {
		x, y = buildLineItem(pdf, x, y, lineItem)
	}

	x, y = width/1.75, y+lineHt*2.25
	x, y = trailerLine(pdf, x, y, "Subtotal", subtotal)
	x, y = trailerLine(pdf, x, y, "Tax", tax)
	pdf.SetDrawColor(180, 180, 180)
	pdf.Line(x+20.0, y, x+220.0, y)
	y = y + lineHt*0.5
	x, y = trailerLine(pdf, x, y, "Total", total)

	if err := pdf.OutputFileAndClose("test.pdf"); err != nil {
		log.Fatal(err)
	}
}

func trailerLine(pdf *gofpdf.Fpdf, x, y float64, label string, amount int) (float64, float64) {
	origX := x
	width, _ := pdf.GetPageSize()
	pdf.SetFont("times", "", 14)
	pdf.SetTextColor(180, 180, 180)
	_, lineHt := pdf.GetFontSize()
	pdf.MoveTo(x, y)
	pdf.CellFormat(80, lineHt, label, gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = width - xIndent - 2.0 - 120.0
	pdf.SetTextColor(50, 50, 50)
	pdf.MoveTo(x, y)
	pdf.CellFormat(120, lineHt, toUSD(amount), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	y = y + lineHt*1.5
	return origX, y
}

func toUSD(cents int) string {
	centsStr := fmt.Sprintf("%d", cents%100)
	if len(centsStr) < 2 {
		centsStr = "0" + centsStr
	}
	return fmt.Sprintf("$%d.%s", cents/100, centsStr)
}

func buildLineItem(pdf *gofpdf.Fpdf, x, y float64, item LineItem) (float64, float64) {
	origX := x
	width, _ := pdf.GetPageSize()
	pdf.SetFont("times", "", 14)
	_, lineHt := pdf.GetFontSize()
	pdf.SetTextColor(50, 50, 50)
	x, y = xIndent-2.0, y+lineHt*.5
	pdf.MoveTo(x, y)
	pdf.MultiCell(width/2.65+1.0, lineHt, item.UnitName, gofpdf.BorderNone, gofpdf.AlignLeft, false)
	tmp := pdf.SplitLines([]byte(item.UnitName), width/2.65+1.0)
	maxY := y + float64(len(tmp))*lineHt - lineHt
	x = x + width/2.25 - xIndent
	pdf.MoveTo(x, y)
	pdf.CellFormat(100, lineHt, toUSD(item.PricePerUnit), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = x + 100.0
	pdf.MoveTo(x, y)
	pdf.CellFormat(80, lineHt, fmt.Sprintf("%d", item.UnitsPurchased), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = width - xIndent - 2.0 - 120.0
	pdf.MoveTo(x, y)
	pdf.CellFormat(120, lineHt, toUSD(item.UnitsPurchased*item.PricePerUnit), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	if maxY > y {
		y = maxY
	}
	y = y + lineHt*1.75
	pdf.SetDrawColor(180, 180, 180)
	pdf.Line(xIndent-10.0, y, width-xIndent+10.0, y)
	return origX, y
}

func summaryBlock(pdf *gofpdf.Fpdf, x, y float64, title string, data ...string) (float64, float64) {
	pdf.SetFont("times", "", 14)
	pdf.SetTextColor(180, 180, 180)
	_, lineHt := pdf.GetFontSize()
	y = y + lineHt
	pdf.Text(x, y, title)
	pdf.SetTextColor(50, 50, 50)
	y = y + lineHt*.25
	for _, str := range data {
		y = y + lineHt*1.25
		pdf.Text(x, y, str)
	}
	return x, y
}
