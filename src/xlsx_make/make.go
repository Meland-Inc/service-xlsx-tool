package xlsx_make

import (
	"fmt"
	"time"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/excel"
	"github.com/Meland-Inc/service-xlsx-tool/src/common/serviceLog"
	"github.com/Meland-Inc/service-xlsx-tool/src/xlsx_make/data"
)

func MakeXlsxTable() {
	path := "bin/output"

	MakeTaskXlsxTable(path)
}

func MakeTaskXlsxTable(path string) {
	serviceLog.Debug("----------- make task xlsx begin ----------------")
	MakeTaskXlsxData()

	sheetName := "task"
	pathFile := fmt.Sprintf("%s/Task_%d.xlsx", path, time.Now().Unix())
	f, err := excel.CreateXlsxFile(pathFile, sheetName, data.TaskXlsxFileHeader)
	if err != nil {
		panic(err)
	}

	dataIndexOffset := len(data.TaskXlsxFileHeader)
	for idx, row := range taskTableRows {
		rowStringList, err1 := row.ToStringList()
		if err != nil {
			err = err1
			serviceLog.Error(err.Error())
			continue
		}
		rowId := fmt.Sprintf("A%d", idx+1+dataIndexOffset)
		f.SetSheetRow(sheetName, rowId, &rowStringList)
	}
	if err != nil {
		panic(err)
	}

	excel.SaveXlsxFile(f, pathFile)
	serviceLog.Debug("--- make task xlsx end, file name[%v] ---", pathFile)
}
