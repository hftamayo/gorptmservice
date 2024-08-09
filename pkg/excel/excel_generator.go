package excel

import (
	"log"

	"github.com/xuri/excelize/v2"
)

type ExcelGenerator struct {
	excel *excelize.File
}

func NewExcelGenerator() *ExcelGenerator {
	excel := excelize.NewFile()
	return &ExcelGenerator{excel: excel}
}

func (eg *ExcelGenerator) GenerateSampleExcel(filename string) error {
	index := eg.excel.NewSheet("Sheet1")
	eg.excel.SetCellValue("Sheet1", "A1", "Hello")
	eg.excel.SetCellValue("Sheet1", "B1", "World")
	eg.excel.SetActiveSheet(index)
	err := eg.excel.SaveAs(filename)
	if err != nil {
		log.Printf("Error generating Excel: %v", err)
		return err
	}
	return nil
}
