package xlsx_export

import (
	"errors"
	"fmt"
	"time"

	xlsxTable "game-message-core/xlsxTableData"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/excel"
	"github.com/Meland-Inc/service-xlsx-tool/src/common/serviceLog"
	"gorm.io/gorm"
)

// TaskList.xlsx

var taskListTableRows = make(map[int32]xlsxTable.TaskListTableRow)

func ParseTaskListParams(v interface{}) (include *xlsxTable.TaskListTableIncludeTasks, err error) {
	iss, ok := v.([][]int)
	if !ok {
		return nil, errors.New("invalid task list parameters")
	}
	include = &xlsxTable.TaskListTableIncludeTasks{}
	for _, is := range iss {
		if len(is) < 2 {
			return include, fmt.Errorf("invalid task list parameters data")
		}
		param := xlsxTable.TaskListTableParam{
			TaskId: int32(is[0]),
			Chance: int32(is[1]),
		}
		include.Param = append(include.Param, param)
		include.ChanceSum += param.Chance
	}
	return
}

func ParseTaskList(rows []map[string]interface{}) (err error) {
	for _, row := range rows {
		if row["id"] == "" || row["id"] == "0" {
			continue
		}
		setting := xlsxTable.TaskListTableRow{
			Id:            excel.IntToInt32(row["id"]),
			Level:         excel.IntToInt32(row["level"]),
			System:        excel.IntToInt32(row["system"]),
			RewardExp:     excel.IntToInt32(row["expReward"]),
			RewardId:      excel.IntToInt32(row["itemReward"]),
			ProgressReset: excel.IntToInt32(row["progressReset"]) == 1,
			NeedMELD:      excel.IntToInt32(row["costMELD"]),
		}

		include, err := ParseTaskListParams(row["includeTask"])
		if err == nil {
			for _, v := range include.Param {
				if _, exist := taskTableRows[v.TaskId]; !exist {
					err = fmt.Errorf(" taskList.xlsx Id[%v], include invalid task id [%v]  ", setting.Id, v.TaskId)
					serviceLog.Error(err.Error())
				}
			}
			if len(include.Param) > 0 {
				setting.SetIncludeTask(include)
			}
		} else {
			err = fmt.Errorf(" taskList.xlsx invalid item taskId[%v]", setting.Id)
			serviceLog.Error(err.Error())
			continue
		}

		taskListTableRows[setting.Id] = setting
	}

	return err
}

func TaskListSaveToDB(db *gorm.DB, curSecUtc time.Time) {
	list := []xlsxTable.TaskListTableRow{}
	for _, TaskList := range taskListTableRows {
		TaskList.CreatedAt = curSecUtc
		list = append(list, TaskList)
	}

	WriterToDB(db, curSecUtc, &xlsxTable.TaskListTableRow{}, len(list), list)
}
