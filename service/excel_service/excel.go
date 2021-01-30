package excel_service

import (
	"github.com/deoooo/gin_demo/pkg/export"
	"github.com/deoooo/gin_demo/pkg/file"
	"github.com/tealeg/xlsx"
	"os"
	"strconv"
	"time"
)

type Data struct {
	Name string
	Age  int
}

func GetData() []*Data {
	var data []*Data
	for i := 0; i < 5; i++ {
		data = append(data, &Data{
			Name: "testName",
			Age:  i,
		})
	}
	return data
}

func Export() (string, error) {
	data := GetData()

	f := xlsx.NewFile()
	sheet, err := f.AddSheet("自动导出数据")
	if err != nil {
		return "", err
	}

	titles := []string{"Name", "Age"}
	row := sheet.AddRow()
	var cell *xlsx.Cell
	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title
	}

	for _, v := range data {
		values := []string{
			v.Name,
			strconv.Itoa(v.Age),
		}

		row = sheet.AddRow()
		for _, value := range values {
			cell = row.AddCell()
			cell.Value = value
		}
	}

	now := strconv.Itoa(int(time.Now().Unix()))
	filename := "data-" + now + ".xlsx"
	fullPath := export.GetExcelFullPath() + filename
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	err = file.IsNotExistMkDir(dir + "/" + export.GetExcelFullPath())
	if err != nil {
		return "", err
	}
	err = f.Save(dir + "/" + fullPath)
	if err != nil {
		return "", err
	}
	return filename, err
}
