package xlsx_export

import (
	"fmt"
	"time"

	xlsxTable "game-message-core/xlsxTableData"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/excel"
	"github.com/Meland-Inc/service-xlsx-tool/src/common/serviceLog"
	"gorm.io/gorm"
)

// Task.xlsx

var taskTableRows = make(map[int32]xlsxTable.TaskTableRow)

func ParseTaskRewardItem(rewardId int32) (objs *xlsxTable.TaskObjectList, err error) {
	objs = &xlsxTable.TaskObjectList{}
	if rewardId < 1 {
		return objs, nil
	}
	if drop, exist := dropTableRows[rewardId]; exist {
		dropList, err := drop.GetDropList()
		if err != nil {
			return nil, err
		}
		for _, dropData := range dropList.List {
			objs.ChanceSum += dropData.Possibility
			objs.ParamList = append(objs.ParamList, xlsxTable.TaskObject{
				Param1: dropData.ObjectCid,
				Param2: dropData.Possibility,
				Param3: dropData.Quality,
			})
		}
	}
	return
}

func ParseTaskParams(v interface{}) (objs *xlsxTable.TaskObjectList, err error) {
	iss, ok := v.([][]int)
	if !ok {
		return &xlsxTable.TaskObjectList{}, nil
	}

	objs = &xlsxTable.TaskObjectList{}
	for _, is := range iss {
		if len(is) < 3 {
			return nil, fmt.Errorf("invalid data")
		}
		paramList := xlsxTable.TaskObject{
			Param1: int32(is[0]),
			Param2: int32(is[1]),
			Param3: int32(is[2]),
		}
		objs.ParamList = append(objs.ParamList, paramList)
		objs.ChanceSum += paramList.Param3
	}
	return
}

func ParseTask(rows []map[string]interface{}) (err error) {
	for _, row := range rows {
		if row["id"] == "" || row["id"] == "0" {
			continue
		}
		setting := xlsxTable.TaskTableRow{
			Id:          excel.IntToInt32(row["id"]),
			Level:       excel.IntToInt32(row["level"]),
			Name:        excel.StringToString(row["name"]),
			SubSystem:   excel.IntSliceToJsonStr(row["subSystem"]),
			Kind:        excel.IntToInt32(row["Type"]),
			RequestLand: excel.IntToInt32(row["requestLand"]),
			RewardExp:   excel.IntToInt32(row["expReward"]),
			Difficulty:  excel.IntToInt32(row["difficulty"]),
		}

		if objs, err := ParseTaskParams(row["item"]); err == nil {
			if len(objs.ParamList) > 0 {
				setting.SetNeedItem(objs)
			}
		} else {
			err = fmt.Errorf(" task.xlsx invalid item taskId[%v]", setting.Id)
			serviceLog.Error(err.Error())
			continue
		}

		if objs, err := ParseTaskParams(row["useItem"]); err == nil {
			if len(objs.ParamList) > 0 {
				setting.SetUseItem(objs)
			}
		} else {
			err = fmt.Errorf(" task.xlsx invalid useItem taskId[%v]", setting.Id)
			serviceLog.Error(err.Error())
			continue
		}

		if objs, err := ParseTaskParams(row["monster"]); err == nil {
			if len(objs.ParamList) > 0 {
				setting.SetKillMonster(objs)
			}
		} else {
			err = fmt.Errorf(" task.xlsx invalid monster taskId[%v]", setting.Id)
			serviceLog.Error(err.Error())
			continue
		}

		if objs, err := ParseTaskParams(row["moveTo"]); err == nil {
			if len(objs.ParamList) > 0 {
				setting.SetTargetPos(objs)
			}
		} else {
			err = fmt.Errorf(" task.xlsx invalid moveTo taskId[%v]", setting.Id)
			serviceLog.Error(err.Error())
			continue
		}

		if objs, err := ParseTaskParams(row["quiz"]); err == nil {
			if len(objs.ParamList) > 0 {
				setting.SetQuiz(objs)
			}
		} else {
			err = fmt.Errorf(" task.xlsx invalid quiz taskId[%v]", setting.Id)
			serviceLog.Error(err.Error())
			continue
		}

		if objs, err := ParseTaskRewardItem(excel.IntToInt32(row["itemReward"])); err == nil {
			if len(objs.ParamList) > 0 {
				setting.SetRewardItems(objs)
			}
		} else {
			err = fmt.Errorf(" task.xlsx invalid quiz taskId[%v]", setting.Id)
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
