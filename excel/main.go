


package main

import (
	"fmt"
	"github.com/tealeg/xlsx"

)

func main() {
	write()
}

//read
func read() {
	exceName := "test_write.xlsx"
	xlFile, err := xlsx.OpenFile(exceName)
	if err != nil {
		panic(err)
	}

	for _, sheet := range xlFile.Sheets {
		fmt.Println("Sheet Name ", sheet.Name)
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Println(text)
			}
		}
	}
}

//write
func write() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row, row1, row2 *xlsx.Row
	var cell *xlsx.Cell
	var err error


	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	row.SetHeightCM(1)
	cell = row.AddCell()
	cell.Value = "姓名"
	cell = row.AddCell()
	cell.Value = "年龄"

	row1 = sheet.AddRow()
    row1.SetHeightCM(1)
    cell = row1.AddCell()
    cell.Value = "狗子"
    cell = row1.AddCell()
    cell.Value = "18"

    row2 = sheet.AddRow()
    row2.SetHeightCM(1)
    cell = row2.AddCell()
    cell.Value = "蛋子"
    cell = row2.AddCell()
    cell.Value = "28"

    err = file.Save("test_write.xlsx")
    if err != nil {
        fmt.Printf(err.Error())
    }
}

//update

func update() {
	excelFileName := "test_write.xlsx"
    xlFile, err := xlsx.OpenFile(excelFileName)
    if err != nil {
        panic(err)
    }
    first := xlFile.Sheets[0]
    row := first.AddRow()
    row.SetHeightCM(1)
    cell := row.AddCell()
    cell.Value = "铁锤"
    cell = row.AddCell()
    cell.Value = "99"

    err = xlFile.Save(excelFileName)
    if err != nil {
        panic(err)
    }
}