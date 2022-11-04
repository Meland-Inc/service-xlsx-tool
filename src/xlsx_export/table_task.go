package xlsx_export

import (
	"encoding/json"
	"fmt"
	"time"

	xlsxTable "game-message-core/xlsxTableData"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/excel"
	"github.com/Meland-Inc/service-xlsx-tool/src/common/serviceLog"
	"gorm.io/gorm"
)

// Task.xlsx

var taskTableRows = make(map[int32]xlsxTable.TaskTableRow)

func ParseTaskParams(parameterJson string) (optionList *xlsxTable.TaskTableRowOptionList, err error) {
	xlsxOptions := &xlsxTable.TaskXlsxOptions{}
	err = json.Unmarshal([]byte(parameterJson), xlsxOptions)
	if err != nil {
		return nil, err
	}
	optionList = &xlsxTable.TaskTableRowOptionList{}
	for _, option := range xlsxOptions.Options {
		tableOption := xlsxTable.TaskTableOption{
			OptionType:      option.OptionType,
			RandomExclusive: 10000, // 随机使用万分比
		}
		intArrArr := excel.ParseIntSliceSliceValue(option.Value)
		for _, intArr := range intArrArr {
			if len(intArr) < 1 {
				continue
			}
			parm := xlsxTable.TaskTableOptionParam{}
			for idx, intValue := range intArr {
				switch idx {
				case 0:
					parm.Param1 = int32(intValue)
				case 1:
					parm.Param2 = int32(intValue)
				case 2:
					parm.Param3 = int32(intValue)
				case 3:
					parm.Param4 = int32(intValue)
				case 4:
					parm.Param5 = int32(intValue)
				}
			}
			tableOption.RandList = append(tableOption.RandList, parm)
		}

		optionList.Options = append(optionList.Options, tableOption)
	}

	return optionList, nil
}

func ParseTask(rows []map[string]interface{}) (err error) {
	for _, row := range rows {
		if row["id"] == "" || row["id"] == "0" {
			continue
		}
		setting := xlsxTable.TaskTableRow{
			Id:         excel.IntToInt32(row["id"]),
			Level:      excel.IntToInt32(row["level"]),
			Name:       excel.StringToString(row["name"]),
			SubSystem:  excel.IntSliceToJsonStr(row["subSystem"]),
			RewardId:   excel.IntToInt32(row["itemReward"]),
			RewardExp:  excel.IntToInt32(row["expReward"]),
			Difficulty: excel.IntToInt32(row["difficulty"]),
			NextTaskId: excel.IntToInt32(row["nextTaskId"]),
		}
		parameterJson := excel.StringToString(row["parameter"])
		optionList, err := ParseTaskParams(parameterJson)
		if err != nil {
			err = fmt.Errorf(" Task.xlsx id[%v] parameter 配置错误", setting.Id)
			serviceLog.Error(err.Error())
			continue
		}
		if err = setting.SetOptions(optionList); err != nil {
			err = fmt.Errorf(" Task.xlsx id[%v] parameter 配置错误", setting.Id)
			serviceLog.Error(err.Error())
			continue
		}
		taskTableRows[setting.Id] = setting
	}

	return err
}

func TaskSaveToDB(db *gorm.DB, curSecUtc time.Time) {
	list := []xlsxTable.TaskTableRow{}
	for _, Task := range taskTableRows {
		Task.CreatedAt = curSecUtc
		list = append(list, Task)
	}

	WriterToDB(db, curSecUtc, &xlsxTable.TaskTableRow{}, len(list), list)
}
