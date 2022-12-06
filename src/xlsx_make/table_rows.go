package xlsx_make

import (
	xlsxTable "game-message-core/xlsxTableData"

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

	for i := 0; i < len(taskTableRows)-1; i++ {
		for j := i + 1; j < len(taskTableRows); j++ {
			if taskTableRows[i].Id > taskTableRows[j].Id {
				taskTableRows[i], taskTableRows[j] = taskTableRows[j], taskTableRows[i]
			}
		}
	}

	taskTableRows = append(taskTableRows, xlsxRow)
}

func AddTaskOption(
	id int32,
	designateOptions []xlsxTable.TaskXlsxOption,
	chanceOptions []xlsxTable.TaskXlsxOption,
) {
	row := data.TaskXlsxRow{
		Id:               id,
		DesignateOptions: designateOptions,
		ChanceOptions:    chanceOptions,
	}
	AddTaskXlsxData(row)
}
