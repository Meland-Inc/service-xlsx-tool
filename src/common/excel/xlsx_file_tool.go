package excel

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func CreateXlsxFile(pathFile string, sheet string, fileHeader [][]string) (*excelize.File, error) {
	f := excelize.NewFile()
	// 设置工作簿的默认工作表
	f.SetSheetName(f.GetSheetName(f.GetActiveSheetIndex()), sheet)

	for idx, row := range fileHeader {
		rowId := fmt.Sprintf("A%d", idx+1)
		f.SetSheetRow(sheet, rowId, &row)
	}

	// 根据指定路径保存文件
	err := SaveXlsxFile(f, pathFile)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func SaveXlsxFile(f *excelize.File, path string) error {
	return f.SaveAs(path)
}
