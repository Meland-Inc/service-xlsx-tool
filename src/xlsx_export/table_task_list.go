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

func ParseTaskListTaskPool(v interface{}) (pools *xlsxTable.TaskListTableTaskPool, err error) {
	iss, ok := v.([][]int)
	if !ok {
		return nil, errors.New("invalid task list task pool")
	}

	pools = &xlsxTable.TaskListTableTaskPool{}
	for _, is := range iss {
		if len(is) < 2 {
			return pools, fmt.Errorf("invalid task list parameters data")
		}
		param := xlsxTable.TaskPoolParam{
			TaskId: int32(is[0]),
			Chance: int32(is[1]),
		}
		pools.Param = append(pools.Param, param)
		pools.ChanceSum += param.Chance
	}
	return
}

func ParseTaskListSequence(v interface{}) (seq *xlsxTable.TaskListTableTaskSequence) {
	is := excel.IntSliceToInt32Slice(v)
	seq = &xlsxTable.TaskListTableTaskSequence{Sequence: is}
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

		taskPools, err := ParseTaskListTaskPool(row["taskPool"])
		if err == nil {
			for _, v := range taskPools.Param {
				if _, exist := taskTableRows[v.TaskId]; !exist {
					err = fmt.Errorf(" taskList.xlsx Id[%v], include invalid task id [%v]  ", setting.Id, v.TaskId)
					serviceLog.Error(err.Error())
				}
			}
			if len(taskPools.Param) > 0 {
				setting.SetTaskPool(taskPools)
			}
		} else {
			err = fmt.Errorf(" taskList.xlsx invalid item Id[%v]", setting.Id)
			serviceLog.Error(err.Error())
			continue
		}

		seq := ParseTaskListSequence(row["taskSequence"])
		if len(seq.Sequence) == 0 && len(taskPools.Param) == 0 {
			err = fmt.Errorf(" taskList.xlsx [%d] taskPool && taskSequence is empty", setting.Id)
			serviceLog.Error(err.Error())
			continue
		}
		setting.SetSequence(seq)
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
