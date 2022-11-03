package xlsx_make

import (
	"github.com/Meland-Inc/service-xlsx-tool/src/common/serviceLog"
	"github.com/Meland-Inc/service-xlsx-tool/src/xlsx_make/data"
)

var taskTableRows = []data.TaskXlsxRow{}

func AddTaskXlsxData(xlsxRow data.TaskXlsxRow) {
	for _, row := range taskTableRows {
		if row.Id == xlsxRow.Id {
			serviceLog.Error("taskRowId[%d] exist in table", row.Id)
			return
		}
	}

	taskTableRows = append(taskTableRows, xlsxRow)
}
