package xlsx_export

import (
	"fmt"
	"time"

	xlsxTable "game-message-core/xlsxTableData"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/excel"
	"github.com/Meland-Inc/service-xlsx-tool/src/common/serviceLog"
	"gorm.io/gorm"
)

// TaskList.xlsx

var taskListTableRows = make(map[int32]xlsxTable.TaskListTableRow)

func ParseTaskListParams(v interface{}) (objs *xlsxTable.TaskObjectList, err error) {
	iss, ok := v.([][]int)
	if !ok {
		return &xlsxTable.TaskObjectList{}, nil
	}

	objs = &xlsxTable.TaskObjectList{}
	for _, is := range iss {
		if len(is) < 2 {
			return nil, fmt.Errorf("invalid data")
		}
		paramList := xlsxTable.TaskObject{
			Param1: int32(is[0]),
			Param2: int32(is[1]),
		}
		objs.ParamList = append(objs.ParamList, paramList)
		objs.ChanceSum += paramList.Param2
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
			ProgressReset: excel.IntToInt32(row["progressReset"]) == 1,
			NeedMELD:      excel.IntToInt32(row["costMELD"]),
		}

		if objs, err := ParseTaskListParams(row["includeTask"]); err == nil {
			for _, v := range objs.ParamList {
				if _, exist := taskTableRows[v.Param1]; !exist {
					err = fmt.Errorf(" taskList.xlsx Id[%v], include invalid task id [%v]  ", setting.Id, v.Param1)
					serviceLog.Error(err.Error())
				}
			}
			if len(objs.ParamList) > 0 {
				setting.SetIncludeTask(objs)
			}
		} else {
			err = fmt.Errorf(" taskList.xlsx invalid item taskId[%v]", setting.Id)
			serviceLog.Error(err.Error())
			continue
		}

		if objs, err := ParseTaskRewardItem(excel.IntToInt32(row["itemReward"])); err == nil {
			if len(objs.ParamList) > 0 {
				setting.SetRewardItems(objs)
			}
		} else {
			err = fmt.Errorf(" taskList.xlsx invalid reward taskId[%v]", setting.Id)
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
