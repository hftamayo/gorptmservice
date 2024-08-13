package services

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/hftamayo/gorptmservice/pkg/excel"
	"github.com/hftamayo/gorptmservice/pkg/pdf"
)

type ReportService struct {
	pdfGenerator   *pdf.PDFGenerator
	excelGenerator *excel.ExcelGenerator
}

func NewReportService() *ReportService {
	return &ReportService{
		pdfGenerator:   pdf.NewPDFGenerator(),
		excelGenerator: excel.NewExcelGenerator(),
	}
}

type ReportRequest struct {
	Data   json.RawMessage `json:"data"`
	Format string          `json:"format"`
}

func (rs *ReportService) GenerateReport(req ReportRequest, filename string) error {
	switch req.Format {
	case "pdf":
		return rs.pdfGenerator.GenerateSamplePDF(filename)
	case "excel":
		return rs.excelGenerator.GenerateSampleExcel(filename)
	default:
		return errors.New("unsupported format")
	}
}

func (rs *ReportService) generatePDFReport(data json.RawMessage, outputPath string) error {
	err := rs.pdfGenerator.GenerateSamplePDF(outputPath)
	if err != nil {
		log.Printf("Error generating PDF report: %v", err)
		return err
	}
	return nil
}

func (rs *ReportService) generateExcelReport(data json.RawMessage, outputPath string) error {
	err := rs.excelGenerator.GenerateSampleExcel(outputPath)
	if err != nil {
		log.Printf("Error generating Excel report: %v", err)
		return err
	}
	return nil
}
