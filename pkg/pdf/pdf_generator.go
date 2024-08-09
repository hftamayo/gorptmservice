package pdf

import (
	"log"

	"github.com/jung-kurt/gofpdf"
)

type PDFGenerator struct {
	pdf *gofpdf.Fpdf
}

func NewPDFGenerator() *PDFGenerator {
	pdf := gofpdf.New("P", "mm", "A4", "")
	return &PDFGenerator{pdf: pdf}
}

func (pg *PDFGenerator) GenerateSamplePDF(filename string) error {
	pg.pdf.AddPage()
	pg.pdf.SetFont("Arial", "B", 16)
	pg.pdf.Cell(40, 10, "Hello, World!")
	err := pg.pdf.OutputFileAndClose(filename)
	if err != nil {
		log.Printf("Error generating PDF: %v", err)
		return err
	}
	return nil
}
