package xlsx_export

import (
	"encoding/json"
	"fmt"
	"time"

	"game-message-core/proto"
	xlsxTable "game-message-core/xlsxTableData"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/excel"
	"github.com/Meland-Inc/service-xlsx-tool/src/common/serviceLog"
	"gorm.io/gorm"
)

// Task.xlsx

var taskTableRows = make(map[int32]xlsxTable.TaskTableRow)

func intArrToTaskRowOption(optType proto.TaskOptionType, arr []int, isChance bool) (*xlsxTable.TaskRowOption, error) {
	if len(arr) < 1 || (isChance && len(arr) < 2) {
		return nil, fmt.Errorf("option invalid")
	}

	opt := &xlsxTable.TaskRowOption{OptionType: optType}
	if isChance {
		opt.Chance = int32(arr[len(arr)-1])
	}

	for idx, value := range arr {
		switch idx {
		case 0:
			opt.Param1 = int32(value)
		case 1:
			opt.Param2 = int32(value)
		case 2:
			opt.Param3 = int32(value)
		case 3:
			opt.Param4 = int32(value)
		}
	}

	return opt, nil
}

func ParseTaskOptionsJson(id int32, js string, isChance bool) (optionList *xlsxTable.TaskRowOptionList, err error) {
	if len(js) == 0 {
		return nil, nil
	}

	if id < 1 {
		return nil, fmt.Errorf("invalid task id[%d]", id)
	}

	xlsxOptions := &xlsxTable.TaskXlsxOptions{}
	err = json.Unmarshal([]byte(js), xlsxOptions)
	if err != nil {
		return nil, err
	}

	optionList = &xlsxTable.TaskRowOptionList{}
	for _, option := range xlsxOptions.Options {
		intArrArr := excel.ParseIntSliceSliceValue(option.Value)
		for _, intArr := range intArrArr {
			opt, err := intArrToTaskRowOption(option.OptionType, intArr, isChance)
			if err != nil {
				serviceLog.Error(err.Error())
				continue
			}
			optionList.Options = append(optionList.Options, *opt)
			optionList.ChanceSum += opt.Chance
		}
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
		}

		// 解析designateOptions
		designateOptionsJs := excel.StringToString(row["designateOptions"])
		desOptList, err := ParseTaskOptionsJson(setting.Id, designateOptionsJs, false)
		if err != nil {
			err = fmt.Errorf(" Task.xlsx id[%v] designateOptions 配置错误", setting.Id)
			serviceLog.Error(err.Error())
			continue
		}
		if err = setting.SetDesignateOptions(desOptList); err != nil {
			serviceLog.Error(err.Error())
			continue
		}

		// 解析designateOptions
		chanceOptionsJs := excel.StringToString(row["chanceOptions"])
		chaOptList, err := ParseTaskOptionsJson(setting.Id, chanceOptionsJs, true)
		if err != nil {
			err = fmt.Errorf(" Task.xlsx id[%v] chanceOptions 配置错误", setting.Id)
			serviceLog.Error(err.Error())
			continue
		}
		if err = setting.SetChanceOptions(chaOptList); err != nil {
			serviceLog.Error(err.Error())
			continue
		}

		if err = setting.Check(); err != nil {
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
